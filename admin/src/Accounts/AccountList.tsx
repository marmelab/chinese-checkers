import { DataTable, List, ReferenceManyCount } from "react-admin";

export function AccountList() {
	return (
		<List resource="accounts">
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
