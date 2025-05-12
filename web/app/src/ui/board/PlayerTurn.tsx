import React from "react";
import "./PlayerTurn.css";
import { Game, getCurrentPlayer } from "../../model/game";
import { GamePlayer } from "../../model/game-player";

export function PlayerTurn({ game }: { game: Game }) {
	const currentPlayer = getCurrentPlayer(game);
	return (
		<p className="player-turn">
			Current player:{" "}
			<strong
				className={
					currentPlayer.gamePlayer == GamePlayer.Green ? "green" : "red"
				}
			>
				{currentPlayer.name}
			</strong>
		</p>
	);
}
