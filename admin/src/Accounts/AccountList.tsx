import { Datagrid, List, NumberField, TextField } from "react-admin";

export function AccountList() {
	return (
		<List resource="accounts">
			<Datagrid>
				<NumberField source="id" />
				<TextField source="name" />
				<TextField source="email" />
			</Datagrid>
		</List>
	);
}
