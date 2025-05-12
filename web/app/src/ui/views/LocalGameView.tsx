import React from "react";
import "./GameView.css";
import { PlayerTurn } from "../board/PlayerTurn";
import { setLocalGame, useLocalGameStore } from "../../storage/local-game";
import { PlayableGameBoard } from "../board/PlayableGameBoard";

export function LocalGameView() {
	const localGame = useLocalGameStore();

	return (
		<>
			<header>
				<h1>Local Game</h1>
			</header>
			<main className={"game"}>
				<PlayableGameBoard game={localGame.game} onChange={setLocalGame} />
				<PlayerTurn game={localGame.game} />
			</main>
		</>
	);
}
