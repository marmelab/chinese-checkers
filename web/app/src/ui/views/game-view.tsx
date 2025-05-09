import React, {useEffect} from "react";
import {useParams} from "react-router-dom";
import {usePageTitle} from "../layout";
import {GameBoard} from "../components/games/game-board";
import {Game, getCurrentPlayer} from "../../model/game";
import {GamePlayer} from "../../model/game-player";
import {useFetchGame} from "../../api/games";

/**
 * Game view component.
 */
export function GameView() {
	usePageTitle("Game");

	// Fetch the game with the UUID from the URL.
	const gameUuid = useParams().uuid;
	const fetchedGame = useFetchGame(gameUuid);

	useEffect(() => {
		// Refetch the game every 5 seconds.
		const interval = setInterval(() => fetchedGame.refetch(), 5000);
		return () => clearInterval(interval);
	}, []);

	return (
		<main className={"game"}>
			<PlayerTurn game={fetchedGame.data}/>
			<GameBoard board={fetchedGame.data.board}/>
		</main>
	);
}

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
