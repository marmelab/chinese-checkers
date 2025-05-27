import {
	EmailField,
	ReferenceManyCount,
	ReferenceManyField,
	Show,
	SimpleShowLayout,
	TextField,
	useRecordContext,
} from "react-admin";
import { GamesTable } from "../Games/GamesTable.tsx";

export function AccountShowTitle() {
	const record = useRecordContext();

	return record?.name ?? "Unknown";
}

export function AccountShow() {
	return (
		<Show title={<AccountShowTitle />}>
			<SimpleShowLayout>
				<TextField source="name" />
				<EmailField source="email" />

				<ReferenceManyCount
					label="Games count"
					reference="online_player"
					target="account_id"
				/>

				<ReferenceManyField
					label="Games"
					reference="accounts_games"
					target="account_id"
					sort={{
						field: "created_at",
						order: "DESC",
					}}
				>
					<GamesTable />
				</ReferenceManyField>
			</SimpleShowLayout>
		</Show>
	);
}
