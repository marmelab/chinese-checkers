import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import "./GameView.css";
import "./OnlineGameView.css";
import { GameBoard } from "../board/GameBoard";
import { useFetchGame } from "../../api/games";
import { PlayerTurn } from "../board/PlayerTurn";
import {
	Game,
	getCurrentPlayer,
	getGameGreenPlayer,
	getGameRedPlayer,
	getWinnerPlayer,
	isGameStarted,
} from "../../model/game";
import { ErrorView } from "./ErrorView";
import { getOnlineGamePlayerId } from "../../storage/online-game";
import { PlayableGameBoard } from "../board/PlayableGameBoard";
import { confetti } from "@tsparticles/confetti";
import { useGameLiveUpdate } from "./useGameLiveUpdate";
import { PlayersWinChances } from "../board/PlayersWinChances";
import { GetMoveHint } from "../board/GetMoveHint";
import { resetMovesHint } from "../../storage/moves-hint";

export function OnlineGameView() {
	const gameUuid = useParams().uuid;
	const fetchedGame = useFetchGame(gameUuid);

	useEffect(() => {
		resetMovesHint();
	}, []);

	const [updatedGame, setUpdatedGame] = useState<Game | null>(null);

	useGameLiveUpdate(gameUuid, setUpdatedGame);

	const game = updatedGame ?? fetchedGame?.data;

	const onlineGamePlayerId = getOnlineGamePlayerId(game?.uuid);

	useEffect(() => {
		if (game?.winner && getWinnerPlayer(game)?.uuid == onlineGamePlayerId)
			confetti({
				particleCount: 100,
				spread: 70,
				origin: { y: 0.6 },
			});
	}, [game, onlineGamePlayerId]);

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

	const isPlayable =
		onlineGamePlayerId && getCurrentPlayer(game).uuid == onlineGamePlayerId;

	return (
		<>
			<header>
				<h1>
					<span className="green">
						{getGameGreenPlayer(game)?.name ?? "Green"}
					</span>{" "}
					VS{" "}
					<span className="red">{getGameRedPlayer(game)?.name ?? "Red"}</span>
				</h1>
			</header>
			<main className="online game">
				<PlayersWinChances game={game} />
				{isPlayable ? (
					<PlayableGameBoard game={game} onChange={setUpdatedGame} online />
				) : (
					<GameBoard board={game.board} />
				)}
				<PlayerTurn game={game} playerId={onlineGamePlayerId} />
				{isPlayable && <GetMoveHint game={game} />}
			</main>
		</>
	);
}
