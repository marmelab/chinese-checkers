import { ApiError, fetchApi } from "./api";
import { Account, zAccount } from "../model/account";
import { useSuspenseQuery } from "@tanstack/react-query";
import { getGame } from "./games";

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

export async function refreshAuthentication(): Promise<string> {
	try {
		return (
			await fetchApi(
				(import.meta.env.VITE_SERVER_NAME ?? "") +
					"/api/v1/authentication/refresh",
				{
					method: "POST",
					headers: {
						"Content-Type": "application/json",
					},
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

export async function logout(): Promise<void> {
	await fetchApi(
		(import.meta.env.VITE_SERVER_NAME ?? "") + "/api/v1/authentication/logout",
		{
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
		},
	);
}

export async function authenticatedAccount(): Promise<Account> {
	return zAccount.parse(
		await fetchApi("/api/v1/authentication/account", {
			method: "GET",
		}),
	);
}

export class InvalidCredentialsError extends Error {}

export function useFetchAuthenticatedAccount() {
	return useSuspenseQuery({
		queryKey: ["authenticated-account"],
		queryFn: () => authenticatedAccount(),
	});
}
