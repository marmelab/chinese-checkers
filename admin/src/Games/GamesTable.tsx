import {
	DataTable,
	DateField,
	FunctionField,
	ReferenceField,
	ReferenceManyField,
	ShowButton,
	SingleFieldList,
} from "react-admin";
import { GameStatusField } from "./GameStatusField.tsx";
import { GamePlayer } from "../../../web/app/src/model/game-player.ts";
import { GamePlayerField } from "./GamePlayerField.tsx";

export function GamesTable() {
	return (
		<DataTable hiddenColumns={["uuid"]}>
			<DataTable.Col source="uuid" label="UUID" />

			<DataTable.Col label="Status" source="status">
				<GameStatusField />
			</DataTable.Col>

			<DataTable.Col label="Green player">
				<FunctionField
					render={(gameRecord) => (
						<ReferenceManyField
							reference="online_player"
							target="game_uuid"
							filter={{
								game_player: GamePlayer.Green.valueOf(),
							}}
						>
							<SingleFieldList linkType={false}>
								<ReferenceField
									reference="accounts"
									source="account_id"
									link="show"
								>
									<GamePlayerField
										gameRecord={gameRecord}
										currentPlayer={GamePlayer.Green.valueOf()}
									/>
								</ReferenceField>
							</SingleFieldList>
						</ReferenceManyField>
					)}
				/>
			</DataTable.Col>

			<DataTable.Col label="Red player">
				<FunctionField
					render={(gameRecord) => (
						<ReferenceManyField
							reference="online_player"
							target="game_uuid"
							filter={{
								game_player: GamePlayer.Red.valueOf(),
							}}
						>
							<SingleFieldList linkType={false}>
								<ReferenceField
									reference="accounts"
									source="account_id"
									link="show"
								>
									<GamePlayerField
										gameRecord={gameRecord}
										currentPlayer={GamePlayer.Red.valueOf()}
									/>
								</ReferenceField>
							</SingleFieldList>
						</ReferenceManyField>
					)}
				/>
			</DataTable.Col>

			<DataTable.Col label="Creation" source="created_at">
				<DateField source="created_at" showTime />
			</DataTable.Col>
			<DataTable.Col label="Update" source="updated_at">
				<DateField source="updated_at" showTime />
			</DataTable.Col>

			<DataTable.Col>
				<ShowButton resource="games" />
			</DataTable.Col>
		</DataTable>
	);
}
