import React from "react";
import "./GameBoard.css";
import { Game } from "../../model/game";
import { GameBoardRow } from "./GameBoardRow";
import { MoveState } from "./PlayableGameBoard";
import { MoveOverlays } from "../move/MoveOverlays";
import { useBestMoveHint } from "../../storage/moves-hint";

export function GameBoard({
	board,
	move,
	onCellClick,
}: {
	board: Game["board"];

	/**
	 * The current move to show on the board.
	 */
	move?: MoveState;

	onCellClick?: (rowIndex: number, columnIndex: number) => void;
}) {
	const bestMoveHint = useBestMoveHint();

	return (
		<>
			<table className="game-board">
				<tbody>
					{board.map((row, rowIndex) => (
						<GameBoardRow
							key={rowIndex}
							move={move}
							row={row}
							rowIndex={rowIndex}
							onClick={onCellClick}
						/>
					))}
				</tbody>
			</table>

			<MoveOverlays move={move} />
			{move?.length == 0 && bestMoveHint && (
				<MoveOverlays move={bestMoveHint} hint />
			)}
		</>
	);
}
