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

/**
 * In milliseconds.
 */
const EVENT_SOURCE_FAILURE_RETRY_TIME = 1000;

function startEventSource(
	gameUuid: string,
	setEventSource: (eventSource: EventSource) => void,
	onGameState: (game: Game) => void,
) {
	const eventSource = new EventSource(`/.well-known/mercure?topic=${gameUuid}`);

	eventSource.addEventListener("message", (event) => {
		onGameState(zGame.parse(JSON.parse(event.data)));
	});
	eventSource.addEventListener("error", () => {
		window.setTimeout(
			() => startEventSource(gameUuid, setEventSource, onGameState),
			EVENT_SOURCE_FAILURE_RETRY_TIME,
		);
	});

	setEventSource(eventSource);
}

export function OnlineGameView() {
	const gameUuid = useParams().uuid;
	const fetchedGame = useFetchGame(gameUuid);

	const eventSourceRef = useRef<EventSource | null>(null);
	const [updatedGame, setUpdatedGame] = useState<Game | null>(null);

	useEffect(() => {
		startEventSource(
			gameUuid,
			(eventSource) => (eventSourceRef.current = eventSource),
			setUpdatedGame,
		);

		return () => {
			eventSourceRef.current.close();
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
