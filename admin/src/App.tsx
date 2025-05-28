import { Admin, Resource } from "react-admin";
import { Layout } from "./Layout";
import { AccountList } from "./Accounts/AccountList.tsx";
import { dataProvider } from "./DataProvider.ts";
import { authProvider } from "./AuthProvider.ts";
import { GameList } from "./Games/GameList.tsx";
import { GameShow } from "./Games/GameShow.tsx";
import { AccountShow } from "./Accounts/AccountShow.tsx";
import { AccountEdit } from "./Accounts/AccountEdit.tsx";

export const App = () => (
	<Admin
		layout={Layout}
		dataProvider={dataProvider}
		authProvider={authProvider}
	>
		<Resource
			name="games"
			list={GameList}
			show={GameShow}
			recordRepresentation={(record) => record?.name ?? record.uuid}
		/>
		<Resource
			name="accounts"
			options={{ label: "Players" }}
			list={AccountList}
			show={AccountShow}
			edit={AccountEdit}
		/>
		<Resource name="online_player" recordRepresentation="name" />
	</Admin>
);
