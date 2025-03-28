import handlebars from "handlebars";
import knex from "knex";
import { customAlphabet } from "nanoid";
import { z } from "zod";

import { withGatewayProxy } from "@app/lib/gateway";
import { alphaNumericNanoId } from "@app/lib/nanoid";

import { TGatewayServiceFactory } from "../../gateway/gateway-service";
import { verifyHostInputValidity } from "../dynamic-secret-fns";
import { DynamicSecretSqlDBSchema, SqlProviders, TDynamicProviderFns } from "./models";

const EXTERNAL_REQUEST_TIMEOUT = 10 * 1000;

const generatePassword = (provider: SqlProviders) => {
  // oracle has limit of 48 password length
  const size = provider === SqlProviders.Oracle ? 30 : 48;

  const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.~!*";
  return customAlphabet(charset, 48)(size);
};

const generateUsername = (provider: SqlProviders) => {
  // For oracle, the client assumes everything is upper case when not using quotes around the password
  if (provider === SqlProviders.Oracle) return alphaNumericNanoId(32).toUpperCase();

  return alphaNumericNanoId(32);
};

type TSqlDatabaseProviderDTO = {
  gatewayService: Pick<TGatewayServiceFactory, "fnGetGatewayClientTls">;
};

export const SqlDatabaseProvider = ({ gatewayService }: TSqlDatabaseProviderDTO): TDynamicProviderFns => {
  const validateProviderInputs = async (inputs: unknown) => {
    const providerInputs = await DynamicSecretSqlDBSchema.parseAsync(inputs);
    verifyHostInputValidity(providerInputs.host, Boolean(providerInputs.projectGatewayId));
    return providerInputs;
  };

  const $getClient = async (providerInputs: z.infer<typeof DynamicSecretSqlDBSchema>) => {
    const ssl = providerInputs.ca ? { rejectUnauthorized: false, ca: providerInputs.ca } : undefined;
    const isMsSQLClient = providerInputs.client === SqlProviders.MsSQL;

    const db = knex({
      client: providerInputs.client,
      connection: {
        database: providerInputs.database,
        port: providerInputs.port,
        host: providerInputs.host,
        user: providerInputs.username,
        password: providerInputs.password,
        ssl,
        // @ts-expect-error this is because of knexjs type signature issue. This is directly passed to driver
        // https://github.com/knex/knex/blob/b6507a7129d2b9fafebf5f831494431e64c6a8a0/lib/dialects/mssql/index.js#L66
        // https://github.com/tediousjs/tedious/blob/ebb023ed90969a7ec0e4b036533ad52739d921f7/test/config.ci.ts#L19
        options: isMsSQLClient
          ? {
              trustServerCertificate: !providerInputs.ca,
              cryptoCredentialsDetails: providerInputs.ca ? { ca: providerInputs.ca } : {}
            }
          : undefined
      },
      acquireConnectionTimeout: EXTERNAL_REQUEST_TIMEOUT
    });
    return db;
  };

  const gatewayProxyWrapper = async (
    providerInputs: z.infer<typeof DynamicSecretSqlDBSchema>,
    gatewayCallback: (host: string, port: number) => Promise<void>
  ) => {
    const relayDetails = await gatewayService.fnGetGatewayClientTls(providerInputs.projectGatewayId as string);
    const [relayHost, relayPort] = relayDetails.relayAddress.split(":");
    await withGatewayProxy(
      async (port) => {
        await gatewayCallback("localhost", port);
      },
      {
        targetHost: providerInputs.host,
        targetPort: providerInputs.port,
        relayHost,
        relayPort: Number(relayPort),
        identityId: relayDetails.identityId,
        orgId: relayDetails.orgId,
        tlsOptions: {
          ca: relayDetails.certChain,
          cert: relayDetails.certificate,
          key: relayDetails.privateKey.toString()
        }
      }
    );
  };

  const validateConnection = async (inputs: unknown) => {
    const providerInputs = await validateProviderInputs(inputs);
    let isConnected = false;
    const gatewayCallback = async (host = providerInputs.host, port = providerInputs.port) => {
      const db = await $getClient({ ...providerInputs, port, host });
      // oracle needs from keyword
      const testStatement = providerInputs.client === SqlProviders.Oracle ? "SELECT 1 FROM DUAL" : "SELECT 1";

      isConnected = await db.raw(testStatement).then(() => true);
      await db.destroy();
    };

    if (providerInputs.projectGatewayId) {
      await gatewayProxyWrapper(providerInputs, gatewayCallback);
    } else {
      await gatewayCallback();
    }
    return isConnected;
  };

  const create = async (inputs: unknown, expireAt: number) => {
    const providerInputs = await validateProviderInputs(inputs);
    const username = generateUsername(providerInputs.client);
    const password = generatePassword(providerInputs.client);
    const gatewayCallback = async (host = providerInputs.host, port = providerInputs.port) => {
      const db = await $getClient({ ...providerInputs, port, host });
      try {
        const { database } = providerInputs;
        const expiration = new Date(expireAt).toISOString();

        const creationStatement = handlebars.compile(providerInputs.creationStatement, { noEscape: true })({
          username,
          password,
          expiration,
          database
        });

        const queries = creationStatement.toString().split(";").filter(Boolean);
        await db.transaction(async (tx) => {
          for (const query of queries) {
            // eslint-disable-next-line
            await tx.raw(query);
          }
        });
      } finally {
        await db.destroy();
      }
    };
    if (providerInputs.projectGatewayId) {
      await gatewayProxyWrapper(providerInputs, gatewayCallback);
    } else {
      await gatewayCallback();
    }
    return { entityId: username, data: { DB_USERNAME: username, DB_PASSWORD: password } };
  };

  const revoke = async (inputs: unknown, entityId: string) => {
    const providerInputs = await validateProviderInputs(inputs);
    const username = entityId;
    const { database } = providerInputs;
    const gatewayCallback = async (host = providerInputs.host, port = providerInputs.port) => {
      const db = await $getClient({ ...providerInputs, port, host });
      try {
        const revokeStatement = handlebars.compile(providerInputs.revocationStatement)({ username, database });
        const queries = revokeStatement.toString().split(";").filter(Boolean);
        await db.transaction(async (tx) => {
          for (const query of queries) {
            // eslint-disable-next-line
            await tx.raw(query);
          }
        });
      } finally {
        await db.destroy();
      }
    };
    if (providerInputs.projectGatewayId) {
      await gatewayProxyWrapper(providerInputs, gatewayCallback);
    } else {
      await gatewayCallback();
    }
    return { entityId: username };
  };

  const renew = async (inputs: unknown, entityId: string, expireAt: number) => {
    const providerInputs = await validateProviderInputs(inputs);
    if (!providerInputs.renewStatement) return { entityId };

    const gatewayCallback = async (host = providerInputs.host, port = providerInputs.port) => {
      const db = await $getClient({ ...providerInputs, port, host });
      const expiration = new Date(expireAt).toISOString();
      const { database } = providerInputs;

      const renewStatement = handlebars.compile(providerInputs.renewStatement)({
        username: entityId,
        expiration,
        database
      });
      try {
        if (renewStatement) {
          const queries = renewStatement.toString().split(";").filter(Boolean);
          await db.transaction(async (tx) => {
            for (const query of queries) {
              // eslint-disable-next-line
              await tx.raw(query);
            }
          });
        }
      } finally {
        await db.destroy();
      }
    };
    if (providerInputs.projectGatewayId) {
      await gatewayProxyWrapper(providerInputs, gatewayCallback);
    } else {
      await gatewayCallback();
    }
    return { entityId };
  };

  return {
    validateProviderInputs,
    validateConnection,
    create,
    revoke,
    renew
  };
};
