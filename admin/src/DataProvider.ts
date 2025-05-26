import postgrestProvider, {
	defaultSchema,
} from "@raphiniert/ra-data-postgrest";
import { fetchUtils } from "react-admin";

export const dataProvider = postgrestProvider({
	apiUrl: import.meta.env.VITE_POSTGREST_URL,
	httpClient: fetchUtils.fetchJson,
	defaultListOp: "eq",
	primaryKeys: new Map([
		["games", ["uuid"]],
		["accounts_games", ["uuid"]],
		["online_player", ["uuid"]],
	]),
	schema: defaultSchema,
});
