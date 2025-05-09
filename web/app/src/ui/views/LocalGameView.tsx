import React from "react";
import {usePageTitle} from "../Layout";
import {PlayerTurn} from "../board/PlayerTurn";
import {useLocalGameStore} from "../../storage/local-game";
import {PlayableGameBoard} from "../board/PlayableGameBoard";

/**
 * Local game view component.
 */
export function LocalGameView() {
	usePageTitle("Local game");

	const localGame = useLocalGameStore();

	return (
		<main className={"game"}>
			<PlayerTurn game={localGame.game} />
			<PlayableGameBoard game={localGame.game} />
		</main>
	);
}
