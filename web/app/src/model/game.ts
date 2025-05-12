import { z } from "zod";
import { OnlinePlayer, zOnlinePlayer } from "./online-player";
import { GamePlayer, zGamePlayer } from "./game-player";
import { CellContent, zCellContent } from "./cell";

export const zGame = z.object({
	uuid: z.string().uuid().optional(),

	board: z.array(z.array(zCellContent)),
	currentPlayer: zGamePlayer,
	winner: zGamePlayer.nullish(),

	players: z.array(zOnlinePlayer),
});

export type Game = z.infer<typeof zGame>;
export type GameBoard = Game["board"];

export function getGameGreenPlayer(game: Game): OnlinePlayer {
	return game.players.find((player) => player.gamePlayer == GamePlayer.Green);
}

export function getGameRedPlayer(game: Game): OnlinePlayer {
	return game.players.find((player) => player.gamePlayer == GamePlayer.Red);
}

/**
 * Get the current online player.
 */
export function getCurrentPlayer(game: Game): OnlinePlayer {
	const player = game.players.find(
		(player) => player.gamePlayer == game.currentPlayer,
	);
	if (!player) throw new Error("Unknown current player.");
	return player;
}

export function getWinnerPlayer(game: Game): OnlinePlayer | null {
	const player = game.players.find(
		(player) => player.gamePlayer == game.winner,
	);
	return player ?? null;
}

export function isPawnPlayable(
	game: Game,
	rowIndex: number,
	cellIndex: number,
): boolean {
	return (
		game.board?.[rowIndex]?.[cellIndex]?.valueOf() ===
		game.currentPlayer.valueOf()
	);
}

export function isCellPlayable(
	game: Game,
	rowIndex: number,
	cellIndex: number,
): boolean {
	return game.board?.[rowIndex]?.[cellIndex] == CellContent.Empty;
}
