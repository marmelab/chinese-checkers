import { Chip } from "@mui/material";
import { FunctionField } from "react-admin";

export const statusLabel: Record<string, string> = {
	started: "Started",
	pending: "Waiting for player",
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

export const statusChoices: { id: keyof typeof statusLabel; name: string }[] =
	Object.entries(statusLabel).map(([id, name]) => ({ id, name }));

export function GameStatusField({ showWinner }: { showWinner?: boolean }) {
	return (
		<FunctionField
			render={(record) => (
				<Chip
					label={
						showWinner && record.status == "finished" && record.winner_name
							? `Winner: ${record.winner_name}`
							: (statusLabel?.[record.status] ?? "Unknown")
					}
					color={statusLabelColor?.[record.status] ?? undefined}
				/>
			)}
		/>
	);
}
