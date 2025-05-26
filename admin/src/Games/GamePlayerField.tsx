import { ChipField } from "react-admin";
import { Tooltip } from "@mui/material";

export function GamePlayerField({
	gameRecord,
	currentPlayer,
}: {
	gameRecord: { winner: number };
	currentPlayer: number;
}) {
	const hasWinner = !!gameRecord.winner;
	const isWinner = gameRecord.winner == currentPlayer;

	const chip = (
		<ChipField
			source="name"
			color={hasWinner ? (isWinner ? "success" : "warning") : undefined}
			sx={hasWinner ? { color: "white!important" } : undefined}
		/>
	);

	if (hasWinner) {
		return (
			<Tooltip title={isWinner ? "Winner" : "Loser"} placement="top" arrow>
				{chip}
			</Tooltip>
		);
	} else return chip;
}
