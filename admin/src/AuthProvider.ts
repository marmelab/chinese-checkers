import { AuthProvider } from "react-admin";
import { jwtDecode } from "jwt-decode";
import { authenticate } from "../../web/app/src/api/accounts.ts";
import { dataProvider } from "./DataProvider.ts";

export const authProvider: AuthProvider = {
	async login({ username, password }) {
		try {
			const token = await authenticate(username, password);
			localStorage.setItem("authentication", token);

			// Check authentication by running a test request.
			await dataProvider.getList("accounts", {
				meta: { limit: 1 },
				pagination: {
					page: 1,
					perPage: 1,
				},
				filter: {},
			});
		} catch (error) {
			localStorage.removeItem("authentication");
			throw new Error(
				"Invalid credentials. You must be an administrator to sign in.",
			);
		}
	},
	async checkError(error) {
		const status = error.status;
		if (status === 401 || status === 403) {
			localStorage.removeItem("authentication");
			throw new Error();
		}
	},
	async checkAuth() {
		if (!localStorage.getItem("authentication")) {
			throw new Error();
		}
	},
	async logout() {
		localStorage.removeItem("authentication");
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
};
