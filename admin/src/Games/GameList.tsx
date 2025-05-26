import {
	ChipField,
	ColumnsButton,
	DataTable,
	DateField,
	FunctionField,
	List,
	ReferenceField,
	ReferenceManyField,
	ShowButton,
	SingleFieldList,
	TopToolbar,
} from "react-admin";
import { GameStatusField } from "./GameStatusField.tsx";
import { GamePlayer } from "../../../web/app/src/model/game-player.ts";

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

				<DataTable.Col label="Status">
					<GameStatusField />
				</DataTable.Col>

				<DataTable.Col label="Green">
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
										<ChipField
											source="name"
											color={
												gameRecord.winner
													? gameRecord.winner == GamePlayer.Green.valueOf()
														? "success"
														: "warning"
													: undefined
											}
											sx={
												gameRecord.winner
													? { color: "white!important" }
													: undefined
											}
										/>
									</ReferenceField>
								</SingleFieldList>
							</ReferenceManyField>
						)}
					/>
				</DataTable.Col>

				<DataTable.Col label="Red">
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
										<ChipField
											source="name"
											color={
												gameRecord.winner
													? gameRecord.winner == GamePlayer.Red.valueOf()
														? "success"
														: "warning"
													: undefined
											}
											sx={
												gameRecord.winner
													? { color: "white!important" }
													: undefined
											}
										/>
									</ReferenceField>
								</SingleFieldList>
							</ReferenceManyField>
						)}
					/>
				</DataTable.Col>

				<DataTable.Col label="Creation">
					<DateField source="created_at" showTime />
				</DataTable.Col>
				<DataTable.Col label="Update">
					<DateField source="updated_at" showTime />
				</DataTable.Col>

				<DataTable.Col>
					<ShowButton />
				</DataTable.Col>
			</DataTable>
		</List>
	);
}
