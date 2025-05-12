import { create } from "zustand";
import { Game } from "../model/game";
import { persist } from "zustand/middleware";
import { GamePlayer, randomGamePlayer } from "../model/game-player";
import { v4 as uuidv4 } from "uuid";

interface LocalGame {
	game: Game;
}

export function getDefaultGame(): Game {
	return {
		uuid: uuidv4(),
		currentPlayer: randomGamePlayer(),
		board: [
			[1, 1, 1, 1, 0, 0, 0],
			[1, 1, 1, 0, 0, 0, 0],
			[1, 1, 0, 0, 0, 0, 0],
			[1, 0, 0, 0, 0, 0, 2],
			[0, 0, 0, 0, 0, 2, 2],
			[0, 0, 0, 0, 2, 2, 2],
			[0, 0, 0, 2, 2, 2, 2],
		],
		players: [
			{
				uuid: uuidv4(),
				name: "Green",
				gamePlayer: GamePlayer.Green,
			},
			{
				uuid: uuidv4(),
				name: "Red",
				gamePlayer: GamePlayer.Red,
			},
		],
	};
}

/**
 * A local game store, stored in local storage.
 */
export const useLocalGameStore = create<LocalGame>()(
	persist(
		() => ({
			game: getDefaultGame(),
		}),
		{
			name: "local-game",
		},
	),
);

export function setLocalGame(game: Game): void {
	useLocalGameStore.setState({
		game,
	});
}
