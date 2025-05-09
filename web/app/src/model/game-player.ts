import {z} from "zod";

/**
 * Game player team.
 */
export enum GamePlayer {
	Green = 1,
	Red = 2,
}

export const zGamePlayer = z.nativeEnum(GamePlayer);

/**
 * Get a random game player.
 */
export function randomGamePlayer(): GamePlayer {
	return Math.round(Math.random()) + 1;
}
