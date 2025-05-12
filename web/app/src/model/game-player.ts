import {z} from "zod";

/**
 * Game player team.
 */
export enum GamePlayer {
	Green = 1,
	Red = 2,
}

export const zGamePlayer = z.nativeEnum(GamePlayer);
