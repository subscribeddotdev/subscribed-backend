import { apiClients } from "@@/common/libs/backendapi/browser";
import { Environment } from "@@/common/libs/backendapi/client";
import { LAST_CHOSEN_ENVIRONMENT } from "@@/constants";
import { Callout, Select } from "@radix-ui/themes";
import { useRouter } from "next/router";
import { useCallback, useEffect, useState } from "react";
import styles from "./EnvironmentSelector.module.css";

interface Props {}

export function EnvironmentSelector({}: Props) {
  const router = useRouter();
  const [error, setError] = useState<unknown>();
  const [envs, setEnvs] = useState<Environment[]>([]);
  const [isLoading, setIsLoading] = useState(false);

  const onEnvChange = useCallback(
    async (newEnv: string) => {
      localStorage.setItem(LAST_CHOSEN_ENVIRONMENT, newEnv);
      await router.push(router.asPath.replace(router.query.environment as string, newEnv));
    },
    [router],
  );

  useEffect(() => {
    (async () => {
      try {
        setIsLoading(true);
        const res = await apiClients().Environments.getEnvironments();
        setEnvs(res.data.data);
      } catch (error) {
        setError(error);
      } finally {
        setIsLoading(false);
      }
    })();
  }, []);

  return (
    <>
      {isLoading && <span hidden data-testid="IsLoading"></span>}
      <Select.Root
        data-loading={isLoading}
        defaultValue={router.query.environment as string}
        onValueChange={onEnvChange}
      >
        <Select.Trigger data-testid="EnvSelector_Trigger" className={styles.trigger} />
        <Select.Content>
          <Select.Group>
            <Select.Label>Environments</Select.Label>
            {envs.map((env) => (
              <Select.Item key={env.id} value={env.id}>
                {env.name}
              </Select.Item>
            ))}
          </Select.Group>
        </Select.Content>
      </Select.Root>

      {error && (
        <Callout.Root color="red" mt="2">
          <Callout.Text>Unable to fetch environments. Please refresh the page.</Callout.Text>
        </Callout.Root>
      )}
    </>
  );
}