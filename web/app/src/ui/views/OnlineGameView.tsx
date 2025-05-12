import React, { useEffect } from "react";
import { useParams } from "react-router-dom";
import "./GameView.css";
import { GameBoard } from "../board/GameBoard";
import { useFetchGame } from "../../api/games";
import { PlayerTurn } from "../board/PlayerTurn";
import { getGameGreenPlayer, getGameRedPlayer } from "../../model/game";

export function OnlineGameView() {
	const gameUuid = useParams().uuid;
	const fetchedGame = useFetchGame(gameUuid);

	useEffect(() => {
		// Refetch the game every 5 seconds.
		const interval = setInterval(() => fetchedGame.refetch(), 5000);
		return () => clearInterval(interval);
	}, []);

	return (
		<>
			<header>
				<h1>
					{getGameGreenPlayer(fetchedGame.data)?.name ?? "Green"} VS{" "}
					{getGameRedPlayer(fetchedGame.data)?.name ?? "Red"}
				</h1>
			</header>
			<main className="game">
				<GameBoard board={fetchedGame.data.board} />
				<PlayerTurn game={fetchedGame.data} />
			</main>
		</>
	);
}
