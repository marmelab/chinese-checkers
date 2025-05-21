import {
	DataTable,
	DateField,
	Labeled,
	ReferenceManyField,
	Show,
	SimpleShowLayout,
} from "react-admin";
import { GameStatusField } from "./GameStatusField.tsx";
import { PlayerTeamField } from "./PlayerTeamField.tsx";

export function GameShow() {
	return (
		<Show
			queryOptions={{
				meta: {
					columns: ["*", "status", "winner_name", "name"],
				},
			}}
		>
			<SimpleShowLayout>
				<DateField source="created_at" showTime />
				<DateField source="updated_at" showTime />

				<Labeled label="Status">
					<GameStatusField />
				</Labeled>

				<ReferenceManyField
					label="Players"
					reference="online_player"
					target="game_uuid"
					sort={{
						field: "game_player",
						order: "ASC",
					}}
				>
					<DataTable bulkActionButtons={false}>
						<DataTable.Col source="name" />
						<DataTable.Col source="game_player">
							<PlayerTeamField />
						</DataTable.Col>
					</DataTable>
				</ReferenceManyField>
			</SimpleShowLayout>
		</Show>
	);
}
