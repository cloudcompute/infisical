import { faBan, faCheck, faCopy } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import { EmptyState, IconButton, Spinner, Tooltip } from "@app/components/v2";
import { useTimedReset } from "@app/hooks";
import {
  useGetIdentityUniversalAuth,
  useGetIdentityUniversalAuthClientSecrets
} from "@app/hooks/api";

import { IdentityAuthFieldDisplay } from "./IdentityAuthFieldDisplay";
import { IdentityUniversalAuthClientSecretsTable } from "./IdentityUniversalAuthClientSecretsTable";
import { ViewAuthMethodProps } from "./types";
import { ViewIdentityContentWrapper } from "./ViewIdentityContentWrapper";

export const ViewIdentityUniversalAuthContent = ({
  identityId,
  onEdit,
  onDelete
}: ViewAuthMethodProps) => {
  const { data, isPending } = useGetIdentityUniversalAuth(identityId);
  const { data: clientSecrets = [], isPending: clientSecretsPending } =
    useGetIdentityUniversalAuthClientSecrets(identityId);

  const [copyTextClientId, isCopyingClientId, setCopyTextClientId] = useTimedReset<string>({
    initialState: "Copy Client ID to clipboard"
  });

  if (isPending || clientSecretsPending) {
    return (
      <div className="flex w-full items-center justify-center">
        <Spinner className="text-mineshaft-400" />
      </div>
    );
  }

  if (!data) {
    return (
      <EmptyState
        icon={faBan}
        title="Could not find Universal Auth associated with this Identity."
      />
    );
  }

  return (
    <ViewIdentityContentWrapper onEdit={onEdit} onDelete={onDelete}>
      <IdentityAuthFieldDisplay label="Access Token TLL (seconds)">
        {data.accessTokenTTL}
      </IdentityAuthFieldDisplay>
      <IdentityAuthFieldDisplay label="Access Token Max TLL (seconds)">
        {data.accessTokenMaxTTL}
      </IdentityAuthFieldDisplay>
      <IdentityAuthFieldDisplay label="Access Token Max Number of Uses">
        {data.accessTokenNumUsesLimit}
      </IdentityAuthFieldDisplay>
      <IdentityAuthFieldDisplay label="Access Token Trusted IPs">
        {data.accessTokenTrustedIps.map((ip) => ip.ipAddress).join(", ")}
      </IdentityAuthFieldDisplay>
      <IdentityAuthFieldDisplay label="Client Secret Trusted IPs">
        {data.clientSecretTrustedIps.map((ip) => ip.ipAddress).join(", ")}
      </IdentityAuthFieldDisplay>
      <div className="col-span-2 my-3">
        <div className="mb-2 border-b border-mineshaft-500">
          <span className="text-sm text-bunker-300">Client ID</span>
        </div>
        <div className="flex items-center gap-2">
          <span className="text-sm">{data.clientId}</span>
          <Tooltip content={copyTextClientId}>
            <IconButton
              ariaLabel="copy icon"
              variant="plain"
              onClick={() => {
                navigator.clipboard.writeText(data.clientId);
                setCopyTextClientId("Copied");
              }}
            >
              <FontAwesomeIcon icon={isCopyingClientId ? faCheck : faCopy} />
            </IconButton>
          </Tooltip>
        </div>
      </div>
      <IdentityUniversalAuthClientSecretsTable
        clientSecrets={clientSecrets}
        identityId={identityId}
      />
    </ViewIdentityContentWrapper>
  );
};
