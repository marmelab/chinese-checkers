import {
	ColumnsButton,
	DataTable,
	EditButton,
	List,
	ReferenceManyCount,
	ShowButton,
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
		<List
			resource="accounts"
			title="Players"
			actions={<AccountListActions />}
			sort={{
				field: "id",
				order: "DESC",
			}}
		>
			<DataTable hiddenColumns={["id"]} bulkActionButtons={false}>
				<DataTable.NumberCol source="id" />
				<DataTable.Col source="name" />
				<DataTable.Col source="email" />
				<DataTable.Col label="Games">
					<ReferenceManyCount
						label="Games"
						reference="online_player"
						target="account_id"
					/>
				</DataTable.Col>
				<DataTable.Col>
					<ShowButton />
					<EditButton />
				</DataTable.Col>
			</DataTable>
		</List>
	);
}
