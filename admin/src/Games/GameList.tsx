import {
	AutocompleteInput,
	ColumnsButton,
	DataTable,
	DateField,
	FilterButton,
	FunctionField,
	List,
	RadioButtonGroupInput,
	ReferenceField,
	ReferenceInput,
	ReferenceManyField,
	ShowButton,
	SingleFieldList,
	TopToolbar,
} from "react-admin";
import { GameStatusField, statusChoices } from "./GameStatusField.tsx";
import { GamePlayer } from "../../../web/app/src/model/game-player.ts";
import { GamePlayerField } from "./GamePlayerField.tsx";

export function GameListActions() {
	return (
		<TopToolbar>
			<FilterButton />
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
				field: "created_at",
				order: "DESC",
			}}
			filters={[
				<RadioButtonGroupInput source="status" choices={statusChoices} />,
				<ReferenceInput
					name="account_id"
					label="Player"
					source="online_player.account_id"
					reference="accounts"
				>
					<AutocompleteInput label="Player" />
				</ReferenceInput>,
			]}
			queryOptions={{
				meta: {
					columns: ["*", "status", "winner_name"],
					embed: ["online_player!inner"],
				},
			}}
		>
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
					<ShowButton />
				</DataTable.Col>
			</DataTable>
		</List>
	);
}
