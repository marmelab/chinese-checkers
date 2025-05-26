import {
	AutocompleteInput,
	ColumnsButton,
	FilterButton,
	List,
	RadioButtonGroupInput,
	ReferenceInput,
	TopToolbar,
} from "react-admin";
import { statusChoices } from "./GameStatusField.tsx";
import { GamesTable } from "./GamesTable.tsx";

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
					<AutocompleteInput
						label="Player"
						sx={{ width: "30ch" }}
						filterToQuery={(query) => ({ "name@ilike": `%${query}%` })}
					/>
				</ReferenceInput>,
			]}
			queryOptions={{
				meta: {
					columns: ["*", "status", "winner_name"],
					embed: ["online_player!inner"],
				},
			}}
		>
			<GamesTable />
		</List>
	);
}
