import React, {useCallback, useState} from "react";
import {Game, isCellPlayable, isPawnPlayable} from "../../model/game";
import {GameBoard} from "./GameBoard";
import {MoveActionsBar} from "../move/MoveActionsBar";

export interface CellIdentifier {
	rowIndex: number;
	cellIndex: number;
}

export type MoveState = CellIdentifier[];

export function PlayableGameBoard({game}: {game: Game}) {
	const [move, setMove] = useState<MoveState>([]);

	/**
	 * Append a cell to the move state.
	 * @param rowIndex Row index of the cell in the board.
	 * @param cellIndex Cell index of the cell in the row.
	 */
	const appendCellToMove = (rowIndex: number, cellIndex: number) => {
		setMove([...move, {rowIndex, cellIndex}]);
	};

	const isMoveStarted = move.length > 0;

	return (
		<>
			<GameBoard
				board={game.board}
				move={move}
				onCellClick={(rowIndex, cellIndex) => {
					if (!isMoveStarted) {
						// Starting the move: append the cell to the move if there is a pawn of the current player on the cell.
						if (isPawnPlayable(game, rowIndex, cellIndex))
							appendCellToMove(rowIndex, cellIndex);
						else alert("You must play a pawn of your color.");
					} else {
						// Continuing the move: append the cell to the move if there is no pawn on the cell.
						if (isCellPlayable(game, rowIndex, cellIndex))
							appendCellToMove(rowIndex, cellIndex);
						else alert("You must move your pawn on a free cell.");
					}
				}}
			/>

			{isMoveStarted && (
				<MoveActionsBar
					move={move}
					onCancel={() => {
						setMove([]);
					}}
					onSubmit={() => {}}
				/>
			)}
		</>
	);
}
