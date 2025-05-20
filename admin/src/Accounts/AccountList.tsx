import { Datagrid, List, TextField } from "react-admin";

export function AccountList() {
	return (
		<List resource="accounts">
			<Datagrid>
				<TextField source="name" />
				<TextField source="email" />
			</Datagrid>
		</List>
	);
}
