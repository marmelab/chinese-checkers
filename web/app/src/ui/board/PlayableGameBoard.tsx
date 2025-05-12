import React, { useState } from "react";
import { Game, isCellPlayable, isPawnPlayable } from "../../model/game";
import { GameBoard } from "./GameBoard";
import { MoveActionsBar } from "../move/MoveActionsBar";
import { openModal } from "../kit/Modals";
import { AlertModal } from "../kit/AlertModal";
import { executeMove } from "../../api/games";
import { getCellName } from "../../model/cell";
import { ApiError } from "../../api/api";

export interface CellIdentifier {
	rowIndex: number;
	cellIndex: number;
}

export type MoveState = CellIdentifier[];

export function PlayableGameBoard({
	game,
	onChange,
}: {
	game: Game;
	onChange: (game: Game) => void;
}) {
	const [move, setMove] = useState<MoveState>([]);

	const appendCellToMove = (rowIndex: number, cellIndex: number) => {
		setMove([...move, { rowIndex, cellIndex }]);
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
			});
		} catch (error) {
			if (error instanceof ApiError) {
				openModal(<AlertModal>{await error.getApiMessage()}</AlertModal>);
			} else throw error;
		}
	};

	const isMoveStarted = move.length > 0;

	return (
		<>
			<GameBoard
				board={game.board}
				move={move}
				onCellClick={(rowIndex, cellIndex) => {
					if (!isMoveStarted) {
						if (isPawnPlayable(game, rowIndex, cellIndex))
							appendCellToMove(rowIndex, cellIndex);
						else
							openModal(
								<AlertModal>You must play a pawn of your color.</AlertModal>,
							);
					} else {
						if (isCellPlayable(game, rowIndex, cellIndex))
							appendCellToMove(rowIndex, cellIndex);
						else
							openModal(
								<AlertModal>You must move your pawn a free cell.</AlertModal>,
							);
					}
				}}
			/>

			{isMoveStarted && (
				<MoveActionsBar
					move={move}
					onCancel={() => {
						setMove([]);
					}}
					onSubmit={submitMove}
				/>
			)}
		</>
	);
}
