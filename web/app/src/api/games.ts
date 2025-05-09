import {z} from "zod";
import {Game, zGame} from "../model/game";

/**
 * Get ongoing games.
 */
export async function getOngoingGames(): Promise<Game[]> {
	return z.array(zGame).parse(await (await fetch("/api/v1/games")).json());
}

/**
 * Get a game from its UUID.
 * @param uuid UUID of the game to get.
 */
export async function getGame(uuid: string): Promise<Game> {
	return zGame.parse(await (await fetch(`/api/v1/games/${uuid}`)).json());
}
