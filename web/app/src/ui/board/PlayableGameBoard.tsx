import React, { useState } from "react";
import { Game, isCellPlayable, isPawnPlayable } from "../../model/game";
import { GameBoard } from "./GameBoard";
import { MoveActionsBar } from "../move/MoveActionsBar";
import { executeMove } from "../../api/games";
import { getCellName } from "../../model/cell";
import { ApiError } from "../../api/api";
import { toast } from "react-toastify";

export interface CellIdentifier {
	rowIndex: number;
	cellIndex: number;
}

export type MoveState = CellIdentifier[];

export function formatErrorMessage(errorMessage: string): string {
	return `${errorMessage[0].toUpperCase()}${errorMessage.slice(1)}.`;
}

export function PlayableGameBoard({
	game,
	onChange,
}: {
	game: Game;
	onChange: (game: Game) => void;
}) {
	const [move, setMove] = useState<MoveState>([]);

	const appendCellToMove = async (rowIndex: number, cellIndex: number) => {
		const newMove = [...move, { rowIndex, cellIndex }];
		setMove(newMove); // Optimistic update.

		if (newMove.length >= 2) {
			// Check new move validity.
			try {
				await executeMove(
					game,
					newMove.map((cell) => getCellName(cell.rowIndex, cell.cellIndex)),
				);
			} catch (error) {
				if (error instanceof ApiError) {
					toast.error(formatErrorMessage(await error.getApiMessage()));
					setMove(move);
				} else throw error;
			}
		}
	};

	const handleCellClick = (rowIndex: number, cellIndex: number) => {
		if (!isMoveStarted) {
			if (!isPawnPlayable(game, rowIndex, cellIndex)) {
				toast.error("You must play a pawn of your color.");
				return;
			}

			appendCellToMove(rowIndex, cellIndex);
			return;
		}

		// If the clicked cell is already in the move, remove all cells after it.
		for (const [index, cell] of move.entries()) {
			if (cell.rowIndex == rowIndex && cell.cellIndex == cellIndex) {
				setMove(move.toSpliced(index, move.length - index));
				return;
			}
		}

		if (!isCellPlayable(game, rowIndex, cellIndex)) {
			toast.error("You must move your pawn a free cell.");
			return;
		}
		appendCellToMove(rowIndex, cellIndex);
	};

	const submitMove = async () => {
		setMove([]);
		try {
			const updatedGame = await executeMove(
				game,
				move.map((cell) => getCellName(cell.rowIndex, cell.cellIndex)),
			);

			onChange({
				...game,
				board: updatedGame.board,
				currentPlayer: updatedGame.currentPlayer,
				winner: updatedGame.winner,
			});
		} catch (error) {
			if (error instanceof ApiError) {
				toast.error(formatErrorMessage(await error.getApiMessage()));
			} else throw error;
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
