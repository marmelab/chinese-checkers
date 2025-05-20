import {
	ChipField,
	Datagrid,
	DateField,
	FunctionField,
	List,
	ReferenceField,
	ReferenceManyField,
	SingleFieldList,
	TextField,
} from "react-admin";
import { Chip } from "@mui/material";

export const statusLabel: Record<string, string> = {
	started: "Started",
	pending: "Pending",
	finished: "Finished",
};
export const statusLabelColor: Record<
	keyof typeof statusLabel,
	"info" | "primary" | "success" | undefined
> = {
	pending: undefined,
	started: "primary",
	finished: "success",
};

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

				<FunctionField
					label={"Status"}
					render={(record) => (
						<Chip
							label={statusLabel?.[record.status] ?? "Unknown"}
							color={statusLabelColor?.[record.status] ?? undefined}
						/>
					)}
				/>

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
