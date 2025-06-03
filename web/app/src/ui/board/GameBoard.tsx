import React from "react";
import "./GameBoard.css";
import { Game } from "../../model/game";
import { GameBoardRow } from "./GameBoardRow";
import { MoveState } from "./PlayableGameBoard";
import { MoveOverlays } from "../move/MoveOverlays";
import { useBestMoveHint } from "../../storage/moves-hint";
import { clsx } from "clsx";
import { GamePlayer } from "../../model/game-player";

export function GameBoard({
	board,
	move,
	lastMove,
	currentPlayer,
	onCellClick,
}: {
	board: Game["board"];

	/**
	 * The current move to show on the board.
	 */
	move?: MoveState;

	/**
	 * The last move to show on the board.
	 */
	lastMove?: MoveState;

	currentPlayer?: GamePlayer;

	onCellClick?: (rowIndex: number, columnIndex: number) => void;
}) {
	const bestMoveHint = useBestMoveHint();

	return (
		<div id="game-board-container">
			<table
				className={clsx("game-board", {
					moving: move?.length > 0,
				})}
			>
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
			{move?.length == 0 && lastMove?.length > 0 && (
				<MoveOverlays
					move={lastMove}
					player={
						currentPlayer == GamePlayer.Green
							? GamePlayer.Red
							: GamePlayer.Green
					}
				/>
			)}
		</div>
	);
}
