import {z} from "zod";
import {OnlinePlayer, zOnlinePlayer} from "./online-player";
import {GamePlayer} from "./game-player";

export const zGame = z.object({
	uuid: z.string().uuid(),

	board: z.array(z.array(z.number().int())),
	currentPlayer: z.nativeEnum(GamePlayer),

	players: z.array(zOnlinePlayer),
});

export type Game = z.infer<typeof zGame>;

/**
 * Find the green player in players list of the game.
 * @param game
 */
export function getGameGreenPlayer(game: Game): OnlinePlayer
{
	return game.players.find(player => player.gamePlayer == GamePlayer.Green);
}

/**
 * Find the red player in players list of the game.
 * @param game
 */
export function getGameRedPlayer(game: Game): OnlinePlayer
{
	console.log(game.players);
	return game.players.find(player => player.gamePlayer == GamePlayer.Red);
}
