import { Admin, Resource } from "react-admin";
import { Layout } from "./Layout";
import { AccountList } from "./Accounts/AccountList.tsx";
import { dataProvider } from "./DataProvider.ts";
import { GameList } from "./Games/GameList.tsx";

export const App = () => (
	<Admin layout={Layout} dataProvider={dataProvider}>
		<Resource name="games" list={GameList} />
		<Resource name="accounts" list={AccountList} />
		<Resource name="online_player" recordRepresentation="name" />
	</Admin>
);
