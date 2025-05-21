import {
	ChipField,
	Datagrid,
	DateField,
	List,
	ReferenceField,
	ReferenceManyField,
	SingleFieldList,
	TextField,
} from "react-admin";
import { GameStatusField } from "./GameStatusField.tsx";

export function GameList() {
	return (
		<List
			resource="games"
			sort={{
				field: "updated_at",
				order: "DESC",
			}}
			queryOptions={{
				meta: {
					columns: ["*", "status"],
				},
			}}
		>
			<Datagrid>
				<TextField source="uuid" />
				<DateField source="created_at" showTime />
				<DateField source="updated_at" showTime />

				<GameStatusField />

				<ReferenceManyField
					label="Green"
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

				<ReferenceManyField
					label="Red"
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
			</Datagrid>
		</List>
	);
}
