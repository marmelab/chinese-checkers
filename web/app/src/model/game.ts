import {z} from "zod";
import {OnlinePlayer, zOnlinePlayer} from "./online-player";
import {GamePlayer, zGamePlayer} from "./game-player";
import {CellContent, zCellContent} from "./cell";

export const zGame = z.object({
	uuid: z.string().uuid().optional(),

	board: z.array(z.array(zCellContent)),
	currentPlayer: zGamePlayer,

	players: z.array(zOnlinePlayer),
});

export type Game = z.infer<typeof zGame>;
export type GameBoard = Game["board"];

/**
 * Find the green player in players list of the game.
 * @param game
 */
export function getGameGreenPlayer(game: Game): OnlinePlayer {
	return game.players.find((player) => player.gamePlayer == GamePlayer.Green);
}

/**
 * Find the red player in players list of the game.
 * @param game
 */
export function getGameRedPlayer(game: Game): OnlinePlayer {
	return game.players.find((player) => player.gamePlayer == GamePlayer.Red);
}

/**
 * Get the current online player.
 * @param game The game from which to get the current player.
 */
export function getCurrentPlayer(game: Game): OnlinePlayer {
	const player = game.players.find(
		(player) => player.gamePlayer == game.currentPlayer,
	);
	if (!player) throw new Error("Unknown current player.");
	return player;
}

/**
 * Find out if the provided pawn is playable by the current player.
 */
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

/**
 * Find out if the provided cell is playable by the current player.
 */
export function isCellPlayable(
	game: Game,
	rowIndex: number,
	cellIndex: number,
): boolean {
	return game.board?.[rowIndex]?.[cellIndex]?.valueOf() === 0;
}
