import React, {useEffect} from "react";
import {useParams} from "react-router-dom";
import {usePageTitle} from "../layout";
import {useQuery} from "@tanstack/react-query";
import {getGame} from "../../api/games";
import {NotFoundView} from "./not-found-view";
import {GameBoard} from "../components/games/game-board";
import {Game, getCurrentPlayer} from "../../model/game";
import {GamePlayer} from "../../model/game-player";
import {LoaderView} from "../components/loader";

/**
 * Get a game from its UUID.
 * @param uuid UUID of the game to get.
 */
function useGame(uuid: string) {
	return useQuery({
		queryKey: ["game", uuid],
		queryFn: () => getGame(uuid),
		retry: false,
	});
}

/**
 * Game view component.
 */
export function GameView() {
	usePageTitle("Game");

	// Fetch the game with the UUID from the URL.
	const gameUuid = useParams().uuid;
	const fetchedGame = useGame(gameUuid);

	useEffect(() => {
		// Refetch the game every 5 seconds.
		const interval = setInterval(() => fetchedGame.refetch(), 5000);
		return () => clearInterval(interval);
	}, []);

	if (fetchedGame.isPending)
		return <LoaderView/>;

	if (fetchedGame.isError)
		return <NotFoundView/>;

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
