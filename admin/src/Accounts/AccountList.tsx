import {
	ColumnsButton,
	DataTable,
	List,
	ReferenceManyCount,
	TopToolbar,
} from "react-admin";

export function AccountListActions() {
	return (
		<TopToolbar>
			<ColumnsButton />
		</TopToolbar>
	);
}

export function AccountList() {
	return (
		<List resource="accounts" actions={<AccountListActions />}>
			<DataTable hiddenColumns={["id"]}>
				<DataTable.NumberCol source="id" />
				<DataTable.Col source="name" />
				<DataTable.Col source="email" />
				<DataTable.Col label={"Games"}>
					<ReferenceManyCount
						label="Games"
						reference="online_player"
						target="account_id"
					/>
				</DataTable.Col>
			</DataTable>
		</List>
	);
}
