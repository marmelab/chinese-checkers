import {
	BooleanField,
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
		<Show
			title={<AccountShowTitle />}
			queryOptions={{
				meta: {
					columns: ["*", "admin"],
				},
			}}
		>
			<SimpleShowLayout>
				<TextField source="name" />
				<EmailField source="email" />

				<BooleanField source="admin" />

				<ReferenceManyCount
					label="Games count"
					reference="online_player"
					target="account_id"
				/>

				<ReferenceManyField
					label="Games"
					reference="accounts_games"
					target="account_id"
				>
					<GamesTable />
				</ReferenceManyField>
			</SimpleShowLayout>
		</Show>
	);
}
