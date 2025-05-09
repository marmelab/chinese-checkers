import React from "react";
import "./GameBoard.css";
import {Game} from "../../model/game";
import {GameBoardRow} from "./GameBoardRow";
import {CellIdentifier, MoveState} from "./PlayableGameBoard";
import {MoveOverlay} from "../move/MoveOverlay";

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

	onCellClick?: (rowIndex: number, cellIndex: number) => void;
}) {
	const moveOverlays: React.ReactElement[] = [];
	if (move) {
		// Create move overlays by finding all atomic move pairs (move between a cell and another).
		let previousCell: CellIdentifier = null;
		for (const cell of move) {
			if (previousCell) {
				moveOverlays.push(
					<MoveOverlay
						key={`overlay-${previousCell.rowIndex}-${previousCell.cellIndex}-${cell.rowIndex}-${cell.cellIndex}`}
						from={previousCell}
						to={cell}
					/>,
				);
			}

			previousCell = cell;
		}
	}

	return (
		<>
			<table className={"game-board"}>
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

			{moveOverlays}
		</>
	);
}
