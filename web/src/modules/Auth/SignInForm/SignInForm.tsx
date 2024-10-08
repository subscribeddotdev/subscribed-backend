import { Alert } from "@@/common/components/Alert/Alert";
import { Button } from "@@/common/components/Button/Button";
import { Input } from "@@/common/components/Input/Input";
import {
  apiClients,
  createApiClients,
  getApiError,
} from "@@/common/libs/backendapi/browser";
import { LAST_CHOSEN_ENVIRONMENT, paths } from "@@/constants";
import { getPaths } from "@@/paths";
import { Box, Flex, Heading, Link, Text } from "@radix-ui/themes";
import { useFormik } from "formik";
import { useCallback, useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import * as yup from "yup";
import { storeTokenOnTheClient, storeUserDetails } from "../token";
import styles from "./SignInForm.module.css";

export function SignInForm() {
  const [error, setError] = useState("");
  const [query] = useSearchParams();
  const navigate = useNavigate();
  const showAccountCreatedSuccessMsg = query.get("signup-succeeded") === "1";

  const postSignIn = useCallback(async () => {
    const { data } = await apiClients().Environments.getEnvironments();

    const lastChosenEnvironment =
      localStorage.getItem(LAST_CHOSEN_ENVIRONMENT) || "";
    if (!lastChosenEnvironment) {
      localStorage.setItem(LAST_CHOSEN_ENVIRONMENT, data.data[0].id);
    }

    const url = getPaths(
      lastChosenEnvironment || data.data[0].id,
    ).dashboardHomepage;
    navigate(url);
  }, [navigate]);

  const f = useFormik({
    validationSchema,
    validateOnChange: false,
    initialValues: { email: "", password: "" },
    onSubmit: async (values) => {
      try {
        setError("");

        const { data } = await createApiClients("").Auth.signIn(values);

        // Store relevant data on the browser
        storeTokenOnTheClient(data.token);
        storeUserDetails(data);

        await postSignIn();
      } catch (error) {
        setError(getApiError(error));
      }
    },
  });

  return (
    <form className={styles.root} onSubmit={f.handleSubmit}>
      <Heading mb="4">Sign in</Heading>

      {showAccountCreatedSuccessMsg && (
        <Alert
          data-testid="SignInForm_Alert_AccountCreated"
          mb="4"
          color="green"
        >
          You account has been created successfully.
        </Alert>
      )}

      <Flex direction="column" gap="2">
        <Input
          name="email"
          required
          type="email"
          onChange={f.handleChange}
          label="E-mail"
          error={f.errors.email}
        />
        <Input
          name="password"
          label="Password"
          type="password"
          onChange={f.handleChange}
          error={f.errors.password}
        />
      </Flex>

      <Button
        mt="4"
        mb="4"
        type="submit"
        disabled={f.isSubmitting}
        loading={f.isSubmitting}
      >
        Sign in
      </Button>

      <Box mb="4">
        <Text size="2">
          {"Don't"} have an account?{" "}
          <Link href={paths.signup}>Set up one.</Link>
        </Text>
      </Box>

      {error && <Alert color="red">{error}</Alert>}
    </form>
  );
}

function validationSchema() {
  return yup.object().shape({
    email: yup.string().email().required(),
    password: yup.string().required(),
  });
}
