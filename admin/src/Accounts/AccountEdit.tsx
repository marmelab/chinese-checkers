import {
	CheckboxGroupInput,
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
		<Edit title={<AccountEditTitle />} mutationMode="optimistic">
			<SimpleForm
				toolbar={
					<Toolbar>
						<SaveButton />
					</Toolbar>
				}
			>
				<TextInput source="name" />

				<CheckboxGroupInput
					source="roles"
					choices={[{ id: "ROLE_ADMIN", name: "Admin" }]}
				/>
			</SimpleForm>
		</Edit>
	);
}
