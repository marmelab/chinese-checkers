import { ApiError, fetchApi } from "./api";

export async function createAccount(
	name: string,
	email: string,
	password: string,
): Promise<void> {
	await fetchApi("/api/v1/accounts", {
		method: "POST",
		body: JSON.stringify({ name, email, password }),
	});
}

export async function authenticate(
	username: string,
	password: string,
): Promise<string> {
	try {
		return (
			await fetchApi(
				(import.meta.env.VITE_SERVER_NAME ?? "") + "/api/v1/authentication",
				{
					method: "POST",
					headers: {
						"Content-Type": "application/json",
					},
					body: JSON.stringify({ username, password }),
				},
			)
		)?.token;
	} catch (error) {
		if (
			error instanceof ApiError &&
			error.errorMessage == "Invalid credentials."
		) {
			throw new InvalidCredentialsError();
		} else throw error;
	}
}

export class InvalidCredentialsError extends Error {}
