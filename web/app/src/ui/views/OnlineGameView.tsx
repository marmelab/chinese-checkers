import React, { useEffect, useRef, useState } from "react";
import { useParams } from "react-router-dom";
import "./GameView.css";
import { GameBoard } from "../board/GameBoard";
import { useFetchGame } from "../../api/games";
import { PlayerTurn } from "../board/PlayerTurn";
import {
	Game,
	getGameGreenPlayer,
	getGameRedPlayer,
	zGame,
} from "../../model/game";

export function OnlineGameView() {
	const gameUuid = useParams().uuid;
	const fetchedGame = useFetchGame(gameUuid);

	const [updatedGame, setUpdatedGame] = useState<Game | null>(null);

	useEffect(() => {
		const eventSource = new EventSource(
			`/.well-known/mercure?topic=${gameUuid}`,
		);

		eventSource.addEventListener("message", (event) => {
			setUpdatedGame(zGame.parse(JSON.parse(event.data)));
		});

		return () => {
			eventSource.close();
		};
	}, []);

	const game = updatedGame ?? fetchedGame?.data;

	return (
		<>
			<header>
				<h1>
					{getGameGreenPlayer(game)?.name ?? "Green"} VS{" "}
					{getGameRedPlayer(game)?.name ?? "Red"}
				</h1>
			</header>
			<main className="game">
				<GameBoard board={game.board} />
				<PlayerTurn game={game} />
			</main>
		</>
	);
}
