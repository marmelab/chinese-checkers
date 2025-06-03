import { create } from "zustand";
import { Game } from "../model/game";
import { persist } from "zustand/middleware";
import { GamePlayer, randomGamePlayer } from "../model/game-player";
import { v4 as uuidv4 } from "uuid";

interface BotGame {
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
				name: "You",
				gamePlayer: GamePlayer.Green,
			},
			{
				uuid: uuidv4(),
				name: "Robot",
				gamePlayer: GamePlayer.Red,
			},
		],
	};
}

/**
 * A bot game store, stored in local storage.
 */
export const useBotGameStore = create<BotGame>()(
	persist(
		() => ({
			game: getDefaultGame(),
		}),
		{
			name: "bot-game",
		},
	),
);

export function setBotGame(game: Game): void {
	useBotGameStore.setState({
		game,
	});
}

export function resetBotGame(): void {
	useBotGameStore.setState({
		game: getDefaultGame(),
	});
}
