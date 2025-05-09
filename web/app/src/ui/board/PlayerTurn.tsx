import React from "react";
import {Game, getCurrentPlayer} from "../../model/game";
import {GamePlayer} from "../../model/game-player";


/**
 * Show the current player turn of a game.
 */
export function PlayerTurn({game}: {
	game: Game;
}) {
	const currentPlayer = getCurrentPlayer(game);
	return (
		<p className={"player-turn"}>
			<strong
				className={currentPlayer.gamePlayer == GamePlayer.Green ? "green" : "red"}>{currentPlayer.name}</strong> to play
		</p>
	);
}
