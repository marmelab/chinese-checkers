import React, {useEffect} from "react";
import {useParams} from "react-router-dom";
import {usePageTitle} from "../Layout";
import {GameBoard} from "../board/GameBoard";
import {useFetchGame} from "../../api/games";
import {PlayerTurn} from "../board/PlayerTurn";

export function GameView() {
	usePageTitle("Game");

	const gameUuid = useParams().uuid;
	const fetchedGame = useFetchGame(gameUuid);

	useEffect(() => {
		// Refetch the game every 5 seconds.
		const interval = setInterval(() => fetchedGame.refetch(), 5000);
		return () => clearInterval(interval);
	}, []);

	return (
		<main className={"game"}>
			<PlayerTurn game={fetchedGame.data} />
			<GameBoard board={fetchedGame.data.board} />
		</main>
	);
}
