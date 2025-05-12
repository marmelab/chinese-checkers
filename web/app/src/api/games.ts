import {z} from "zod";
import {Game, zGame} from "../model/game";
import {useSuspenseQuery} from "@tanstack/react-query";
import {CellIdentifier, MoveState} from "../ui/board/PlayableGameBoard";

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

/**
 * Execute the provided move on the provided game state.
 * @param game Game state for the move.
 * @param move The move to execute (all visited cell names).
 */
export async function executeMove(game: Game, move: string[]): Promise<Game> {
	return zGame.parse(
		await (
			await fetch("/api/v1/games/move", {
				method: "POST",
				body: JSON.stringify({
					game: game,
					move: move,
				}),
			})
		).json(),
	);
}

/**
 * Fetch ongoing games.
 */
export function useFetchOngoingGames() {
	return useSuspenseQuery({
		queryKey: ["ongoingGames"],
		queryFn: getOngoingGames,
	});
}

/**
 * Fetch a game from its UUID.
 * @param uuid UUID of the game to get.
 */
export function useFetchGame(uuid: string) {
	return useSuspenseQuery({
		queryKey: ["game", uuid],
		queryFn: () => getGame(uuid),
		retry: false,
	});
}
