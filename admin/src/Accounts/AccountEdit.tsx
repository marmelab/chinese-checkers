import { Edit, SimpleForm, TextInput, useRecordContext } from "react-admin";

export function AccountEditTitle() {
	const record = useRecordContext();

	return "Edit " + (record?.name ?? "Unknown");
}

export function AccountEdit() {
	return (
		<Edit title={<AccountEditTitle />}>
			<SimpleForm>
				<TextInput source="name" />
			</SimpleForm>
		</Edit>
	);
}
