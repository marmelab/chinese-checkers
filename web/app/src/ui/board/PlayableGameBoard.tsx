import React, { useState } from "react";
import { toast } from "react-toastify";
import { Game, isCellPlayable, isPawnPlayable } from "../../model/game";
import { GameBoard } from "./GameBoard";
import { MoveActionsBar } from "../move/MoveActionsBar";
import { executeMove, getValidMoves } from "../../api/games";
import { CellIdentifier, getCellName } from "../../model/cell";
import { ApiError } from "../../api/api";
import { showErrorToast } from "../showErrorToast";
import {
	findValidMoveToCell,
	resetMovesHint,
	resetValidMoves,
	setValidMoves,
	useValidMoves,
} from "../../storage/moves-hint";

export type MoveState = CellIdentifier[];

async function fetchMoveState(game: Game, cell: CellIdentifier) {
	try {
		setValidMoves(await getValidMoves(game, cell));
	} catch (error) {
		showErrorToast(error);
	}
}

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

	const validMoves = useValidMoves();

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

			fetchMoveState(game, { row: rowIndex, column: columnIndex });
			appendCellToMove(rowIndex, columnIndex);
			return;
		}

		// If the clicked cell is already in the move, remove all cells after it.
		for (const [index, cell] of move.entries()) {
			if (cell.row == rowIndex && cell.column == columnIndex) {
				const newMove = move.toSpliced(index, move.length - index);
				setMove(newMove);
				if (newMove?.length == 0) resetValidMoves();

				return;
			}
		}

		if (validMoves) {
			const move = findValidMoveToCell(validMoves, {
				row: rowIndex,
				column: columnIndex,
			});
			if (move) {
				setMove(move);
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
			<GameBoard
				board={game.board}
				move={move}
				lastMove={game.lastMove}
				currentPlayer={game.currentPlayer}
				onCellClick={handleCellClick}
			/>

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
