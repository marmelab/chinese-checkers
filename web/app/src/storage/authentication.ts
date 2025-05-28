import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";

export interface AuthenticatedAccount {
	username: string;
	roles: string[];
	exp: number;
}

export function useAuthenticatedAccount(): AuthenticatedAccount | null {
	const authToken = Cookies.get("authentication");

	if (!authToken) return null;

	try {
		const authenticatedAccount = jwtDecode<AuthenticatedAccount>(authToken);

		const now = new Date();
		if (now.getTime() / 1000 >= authenticatedAccount.exp) return null;

		return authenticatedAccount;
	} catch (err) {
		console.error(err);
		return null;
	}
}

export function clearAuthenticatedAccount(): void {
	Cookies.set("authentication", null);
}
