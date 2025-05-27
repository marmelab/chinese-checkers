import {
	ChipField,
	DateField,
	DeleteButton,
	FunctionField,
	Labeled,
	ReferenceManyField,
	Show,
	SimpleShowLayout,
	SingleFieldList,
} from "react-admin";
import { GameStatusField } from "./GameStatusField.tsx";
import { GameBoardField } from "./GameBoard.tsx";
import { Grid } from "@mui/material";

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
				<Grid container justifyContent="end">
					<Grid>
						<DeleteButton />
					</Grid>
				</Grid>
				<Grid container spacing={3}>
					<Grid>
						<Labeled label="Status">
							<GameStatusField showWinner />
						</Labeled>
					</Grid>
					<Grid>
						<Labeled label="Players">
							<ReferenceManyField
								reference="online_player"
								target="game_uuid"
								sort={{
									field: "game_player",
									order: "ASC",
								}}
								queryOptions={{
									meta: {
										embed: ["accounts"],
									},
								}}
							>
								<SingleFieldList linkType={false}>
									<FunctionField
										render={(record) => (
											<ChipField
												source="accounts.name"
												color={record.game_player == 1 ? "success" : "error"}
												variant="outlined"
											/>
										)}
									/>
								</SingleFieldList>
							</ReferenceManyField>
						</Labeled>
					</Grid>
				</Grid>

				<Labeled label="Board">
					<GameBoardField />
				</Labeled>

				<DateField source="created_at" showTime />
				<DateField source="updated_at" showTime />
			</SimpleShowLayout>
		</Show>
	);
}
