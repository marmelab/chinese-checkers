import React, { useState } from "react";
import { toast } from "react-toastify";
import { Game, isCellPlayable, isPawnPlayable } from "../../model/game";
import { GameBoard } from "./GameBoard";
import { MoveActionsBar } from "../move/MoveActionsBar";
import { executeMove } from "../../api/games";
import { CellIdentifier, getCellName } from "../../model/cell";
import { ApiError } from "../../api/api";
import { showErrorToast } from "../showErrorToast";
import { resetMovesHint, useBestMoveHint } from "../../storage/moves-hint";

export type MoveState = CellIdentifier[];

export function PlayableGameBoard({
	game,
	onChange,
	online,
}: {
	game: Game;
	onChange: (game: Game) => void;
	online?: boolean;
}) {
	const [move, setMove] = useState<MoveState>([]);
	online = !!online;

	const appendCellToMove = async (rowIndex: number, columnIndex: number) => {
		const newMove: MoveState = [
			...move,
			{ row: rowIndex, column: columnIndex },
		];
		setMove(newMove); // Optimistic update.

		if (newMove.length >= 2) {
			// Check new move validity.
			try {
				await executeMove(
					game,
					newMove.map((cell) => getCellName(cell.row, cell.column)),
					// Never online as it's a simulated move.
					false,
				);
			} catch (error) {
				showErrorToast(error);
				if (error instanceof ApiError) {
					setMove(move);
				} else throw error;
			}
		}
	};

	const handleCellClick = (rowIndex: number, columnIndex: number) => {
		if (!isMoveStarted) {
			if (!isPawnPlayable(game, rowIndex, columnIndex)) {
				toast.error("You must play a pawn of your color.");
				return;
			}

			appendCellToMove(rowIndex, columnIndex);
			return;
		}

		// If the clicked cell is already in the move, remove all cells after it.
		for (const [index, cell] of move.entries()) {
			if (cell.row == rowIndex && cell.column == columnIndex) {
				setMove(move.toSpliced(index, move.length - index));
				return;
			}
		}

		if (!isCellPlayable(game, rowIndex, columnIndex)) {
			toast.error("You must move your pawn a free cell.");
			return;
		}
		appendCellToMove(rowIndex, columnIndex);
	};

	const submitMove = async () => {
		setMove([]);
		try {
			const updatedGame = await executeMove(
				game,
				move.map((cell) => getCellName(cell.row, cell.column)),
				online,
			);

			resetMovesHint();
			onChange({
				...game,
				board: updatedGame.board,
				currentPlayer: updatedGame.currentPlayer,
				winner: updatedGame.winner,
			});
		} catch (error) {
			showErrorToast(error);
		}
	};

	const isMoveStarted = move.length > 0;

	return (
		<>
			<GameBoard board={game.board} move={move} onCellClick={handleCellClick} />

			{isMoveStarted && (
				<MoveActionsBar
					move={move}
					onCancel={() => setMove([])}
					onSubmit={submitMove}
				/>
			)}
		</>
	);
}
