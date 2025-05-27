import {
	Edit,
	SaveButton,
	SimpleForm,
	TextInput,
	Toolbar,
	useRecordContext,
} from "react-admin";

export function AccountEditTitle() {
	const record = useRecordContext();

	return "Edit " + (record?.name ?? "Unknown");
}

export function AccountEdit() {
	return (
		<Edit title={<AccountEditTitle />}>
			<SimpleForm
				toolbar={
					<Toolbar>
						<SaveButton />
					</Toolbar>
				}
			>
				<TextInput source="name" />
			</SimpleForm>
		</Edit>
	);
}
