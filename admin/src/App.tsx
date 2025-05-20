import { Admin, Resource } from "react-admin";
import { Layout } from "./Layout";
import { AccountList } from "./Accounts/AccountList.tsx";
import { dataProvider } from "./DataProvider.ts";

export const App = () => (
	<Admin layout={Layout} dataProvider={dataProvider}>
		<Resource name="accounts" list={AccountList} />
	</Admin>
);
