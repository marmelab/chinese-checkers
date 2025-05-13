import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import "./GameView.css";
import "./OnlineGameView.css";
import { GameBoard } from "../board/GameBoard";
import { useFetchGame } from "../../api/games";
import { PlayerTurn } from "../board/PlayerTurn";
import {
	Game,
	getGameGreenPlayer,
	getGameRedPlayer,
	isGameStarted,
	zGame,
} from "../../model/game";
import { ErrorView } from "./ErrorView";

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

	if (!game) return <ErrorView />;

	if (!isGameStarted(game))
		return (
			<>
				<header>
					<h1>Online Game</h1>
				</header>
				<main className="online game">
					<p className="center">Waiting for another player to join...</p>

					<p className="join-code">{game.joinCode}</p>

					<p className="center">
						Share this code with your friend so they can join your game.
					</p>
				</main>
			</>
		);

	return (
		<>
			<header>
				<h1>
					{getGameGreenPlayer(game)?.name ?? "Green"} VS{" "}
					{getGameRedPlayer(game)?.name ?? "Red"}
				</h1>
			</header>
			<main className="online game">
				<GameBoard board={game.board} />
				<PlayerTurn game={game} />
			</main>
		</>
	);
}
