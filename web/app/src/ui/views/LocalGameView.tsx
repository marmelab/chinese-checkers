import React, { useEffect } from "react";
import "./GameView.css";
import { PlayerTurn } from "../board/PlayerTurn";
import { setLocalGame, useLocalGameStore } from "../../storage/local-game";
import { PlayableGameBoard } from "../board/PlayableGameBoard";
import { PlayersWinChances } from "../board/PlayersWinChances";
import { GetMoveHint } from "../board/GetMoveHint";
import { resetMovesHint } from "../../storage/moves-hint";

export function LocalGameView() {
	const localGame = useLocalGameStore();

	useEffect(() => {
		resetMovesHint();
	}, []);

	return (
		<>
			<header>
				<h1>Local Game</h1>
			</header>
			<main className={"game"}>
				<PlayersWinChances game={localGame.game} />
				<PlayableGameBoard game={localGame.game} onChange={setLocalGame} />
				<PlayerTurn game={localGame.game} />
				<GetMoveHint game={localGame.game} />
			</main>
		</>
	);
}
