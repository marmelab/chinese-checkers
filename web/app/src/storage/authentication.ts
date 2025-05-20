import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";

export interface AuthenticatedAccount {
	username: string;
	roles: string[];
}

export function useAuthenticatedAccount(): AuthenticatedAccount | null {
	const authToken = Cookies.get("authentication");

	if (!authToken) return null;

	try {
		return jwtDecode<AuthenticatedAccount>(authToken);
	} catch (err) {
		console.error(err);
		return null;
	}
}

export function clearAuthenticatedAccount(): void {
	Cookies.set("authentication", null);
}
