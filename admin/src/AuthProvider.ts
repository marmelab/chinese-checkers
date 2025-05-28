import { addRefreshAuthToAuthProvider, AuthProvider } from "react-admin";
import { jwtDecode } from "jwt-decode";
import {
	authenticate,
	logout,
	refreshAuthentication,
} from "../../web/app/src/api/accounts.ts";

export function checkAndStoreToken(token: string): void {
	if (jwtDecode<any>(token)?.role != "admin")
		throw new Error("You must be an administrator to sign in.");

	localStorage.setItem("authentication", token);
}

export async function tryAuthenticationRefresh() {
	const { exp } = jwtDecode<any>(localStorage.getItem("authentication") ?? "");

	const in15Minutes = new Date();
	in15Minutes.setMinutes(in15Minutes.getMinutes() + 15);

	if (in15Minutes.getTime() / 1000 < exp) return;

	try {
		checkAndStoreToken(await refreshAuthentication());
	} catch (error) {
		await authProvider.logout(undefined);
		throw new Error(
			"Invalid authentication. You must be an administrator to sign in.",
		);
	}
}

export const authProvider: AuthProvider = addRefreshAuthToAuthProvider(
	{
		async login({ username, password }) {
			try {
				checkAndStoreToken(await authenticate(username, password));
			} catch (error) {
				throw new Error(
					"Invalid credentials. You must be an administrator to sign in.",
				);
			}
		},
		async checkError(error) {
			if (
				error.status === 401 ||
				error.status === 403 ||
				error.body?.code === 22023
			) {
				await this.logout();
				throw error;
			}
		},
		async checkAuth() {
			if (!localStorage.getItem("authentication")) {
				throw new Error();
			}
		},
		async logout() {
			localStorage.removeItem("authentication");
			await logout();
		},
		async getIdentity() {
			const { uuid, username, jwtData } = jwtDecode<any>(
				localStorage.getItem("authentication") ?? "",
			);

			return {
				id: uuid,
				fullName: username,
				...jwtData,
			};
		},
		async canAccess() {
			return true;
		},
	},
	tryAuthenticationRefresh,
);
