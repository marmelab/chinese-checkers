import {
	Datagrid,
	List,
	NumberField,
	ReferenceManyCount,
	TextField,
} from "react-admin";

export function AccountList() {
	return (
		<List resource="accounts">
			<Datagrid>
				<NumberField source="id" />
				<TextField source="name" />
				<TextField source="email" />
				<ReferenceManyCount
					label="Games"
					reference="online_player"
					target="account_id"
				/>
			</Datagrid>
		</List>
	);
}
