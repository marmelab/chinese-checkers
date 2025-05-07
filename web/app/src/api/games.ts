import {z} from "zod";
import {Game, zGame} from "../model/game";

/**
 * Get ongoing games.
 */
export async function getOngoingGames(): Promise<Game[]>
{
	return z.array(zGame).parse(await (await fetch("/api/v1/games")).json());
}
