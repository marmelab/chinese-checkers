import { FunctionField } from "react-admin";
// Import colors.
import "../../../web/app/src/ui/Colors.css";
import { GameBoard } from "../../../web/app/src/ui/board/GameBoard.tsx";

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
			render={(record) => <GameBoard board={record.board} />}
		/>
	);
}
