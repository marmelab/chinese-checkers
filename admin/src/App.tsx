import { Admin, Resource } from "react-admin";
import { Layout } from "./Layout";
import { AccountList } from "./Accounts/AccountList.tsx";
import { dataProvider } from "./DataProvider.ts";
import { GameList } from "./Games/GameList.tsx";
import { GameShow } from "./Games/GameShow.tsx";
import { AccountShow } from "./Accounts/AccountShow.tsx";

export const App = () => (
	<Admin layout={Layout} dataProvider={dataProvider}>
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
		/>
		<Resource name="online_player" recordRepresentation="name" />
	</Admin>
);
