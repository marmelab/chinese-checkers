import postgrestProvider, {
	defaultSchema,
} from "@raphiniert/ra-data-postgrest";
import { addRefreshAuthToDataProvider, fetchUtils, Options } from "react-admin";
import { tryAuthenticationRefresh } from "./AuthProvider.ts";

export function fetchJson(
	url: any,
	options: Options,
): ReturnType<typeof fetchUtils.fetchJson> {
	const authToken = localStorage.getItem("authentication");

	return fetchUtils.fetchJson(url, {
		...options,
		user: {
			authenticated: !!authToken,
			token: authToken ? `Bearer ${authToken}` : undefined,
		},
	});
}

export const dataProvider = addRefreshAuthToDataProvider(
	postgrestProvider({
		apiUrl: import.meta.env.VITE_POSTGREST_URL,
		httpClient: fetchJson,
		defaultListOp: "eq",
		primaryKeys: new Map([
			["games", ["uuid"]],
			["accounts_games", ["uuid"]],
			["online_player", ["uuid"]],
		]),
		schema: defaultSchema,
	}),
	tryAuthenticationRefresh,
);
