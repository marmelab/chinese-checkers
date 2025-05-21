import { Chip, Tooltip } from "@mui/material";
import { FunctionField } from "react-admin";

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

export function GameStatusField() {
	return (
		<FunctionField
			label={"Status"}
			render={(record) => {
				const chip = (
					<Chip
						label={statusLabel?.[record.status] ?? "Unknown"}
						color={statusLabelColor?.[record.status] ?? undefined}
					/>
				);

				if (!record.winner_name) return chip;

				return (
					<Tooltip
						title={`Winner: ${record.winner_name}`}
						placement="top"
						arrow
					>
						{chip}
					</Tooltip>
				);
			}}
		/>
	);
}
