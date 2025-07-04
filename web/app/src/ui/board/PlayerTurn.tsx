import React from "react";
import "./PlayerTurn.css";
import { Game, getCurrentPlayer, getWinnerPlayer } from "../../model/game";
import { GamePlayer } from "../../model/game-player";

export function PlayerTurn({
	game,
	playerId,
}: {
	game: Game;
	playerId?: string;
}) {
	const currentPlayer = getCurrentPlayer(game);
	const winnerPlayer = getWinnerPlayer(game);
	return (
		<>
			<p className="player-turn">
				{winnerPlayer ? "Winner" : "Current player"}:{" "}
				<strong
					className={
						(winnerPlayer ?? currentPlayer).gamePlayer == GamePlayer.Green
							? "green"
							: "red"
					}
				>
					{(winnerPlayer ?? currentPlayer).name}
				</strong>
			</p>
			<p className="center">
				{!winnerPlayer &&
					playerId &&
					(currentPlayer.uuid == playerId
						? "Move a pawn!"
						: "Wait for your opponent to move a pawn.")}
			</p>
		</>
	);
}
