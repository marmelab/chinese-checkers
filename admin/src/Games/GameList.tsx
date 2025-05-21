import {
	ChipField,
	ColumnsButton,
	DataTable,
	DateField,
	List,
	ReferenceField,
	ReferenceManyField,
	SingleFieldList,
	TopToolbar,
} from "react-admin";
import { GameStatusField } from "./GameStatusField.tsx";

export function GameListActions() {
	return (
		<TopToolbar>
			<ColumnsButton />
		</TopToolbar>
	);
}

export function GameList() {
	return (
		<List
			resource="games"
			actions={<GameListActions />}
			sort={{
				field: "updated_at",
				order: "DESC",
			}}
			queryOptions={{
				meta: {
					columns: ["*", "status", "winner_name"],
				},
			}}
		>
			<DataTable hiddenColumns={["uuid"]}>
				<DataTable.Col source="uuid" label="UUID" />
				<DataTable.Col label="Creation">
					<DateField source="created_at" showTime />
				</DataTable.Col>
				<DataTable.Col label="Update">
					<DateField source="updated_at" showTime />
				</DataTable.Col>

				<DataTable.Col label="Status">
					<GameStatusField />
				</DataTable.Col>

				<DataTable.Col label="Green">
					<ReferenceManyField
						reference="online_player"
						target="game_uuid"
						filter={{
							game_player: 1,
						}}
					>
						<SingleFieldList linkType={false}>
							<ReferenceField
								reference="accounts"
								source="account_id"
								link={"show"}
							>
								<ChipField source="name" />
							</ReferenceField>
						</SingleFieldList>
					</ReferenceManyField>
				</DataTable.Col>

				<DataTable.Col label="Red">
					<ReferenceManyField
						reference="online_player"
						target="game_uuid"
						filter={{
							game_player: 2,
						}}
					>
						<SingleFieldList linkType={false}>
							<ReferenceField
								reference="accounts"
								source="account_id"
								link={"show"}
							>
								<ChipField source="name" />
							</ReferenceField>
						</SingleFieldList>
					</ReferenceManyField>
				</DataTable.Col>
			</DataTable>
		</List>
	);
}
