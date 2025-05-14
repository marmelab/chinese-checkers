import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";

export interface AuthenticationState {
	username: string;
	roles: string[];
}

export function getAuthenticationState(): AuthenticationState | null {
	const authToken = Cookies.get("authentication");

	if (!authToken) return null;

	try {
		return jwtDecode<AuthenticationState>(authToken);
	} catch (err) {
		console.error(err);
		return null;
	}
}
