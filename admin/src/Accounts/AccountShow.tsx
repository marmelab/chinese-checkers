import {
	EmailField,
	ReferenceManyCount,
	ReferenceManyField,
	Show,
	SimpleShowLayout,
	TextField,
} from "react-admin";
import { GamesTable } from "../Games/GamesTable.tsx";

export function AccountShow() {
	return (
		<Show>
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
				>
					<GamesTable />
				</ReferenceManyField>
			</SimpleShowLayout>
		</Show>
	);
}
