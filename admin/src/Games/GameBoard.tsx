import { FunctionField } from "react-admin";
// Import colors.
import "../../../web/app/src/ui/Colors.css";
import { GameBoard } from "../../../web/app/src/ui/board/GameBoard.tsx";
import { useGameLiveUpdate } from "../../../web/app/src/ui/views/useGameLiveUpdate.ts";
import { useState } from "react";

export function LiveGameBoard({ record }: { record: any }) {
	const [updatedBoard, setUpdatedBoard] = useState<number[][] | null>(null);

	useGameLiveUpdate(
		record.uuid,
		({ board }) => setUpdatedBoard(board),
		import.meta.env.VITE_SERVER_NAME,
	);

	return <GameBoard board={updatedBoard ?? record.board} />;
}

export function GameBoardField() {
	return (
		<FunctionField
			sx={{
				["table.game-board td button"]: {
					width: "3em",
					height: "3em",
					[".pawn"]: {
						width: "2em",
						height: "2em",
					},
				},
			}}
			label="Board"
			render={(record) => <LiveGameBoard record={record} />}
		/>
	);
}
