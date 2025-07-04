import { z } from "zod";
import { Game, zGame } from "../model/game";
import { useSuspenseQuery } from "@tanstack/react-query";
import { fetchApi } from "./api";
import { GameEvaluation, zGameEvaluation } from "../model/game-evaluation";
import { CellIdentifier, getCellName, zCellIdentifier } from "../model/cell";
import { PathsTree, zPathsTree } from "../model/paths";

export async function getOngoingGames(): Promise<Game[]> {
	return z.array(zGame).parse(await fetchApi("/api/v1/games"));
}

/**
 * @param uuid UUID of the game to get.
 */
export async function getGame(uuid: string): Promise<Game> {
	return zGame.parse(await fetchApi(`/api/v1/games/${uuid}`));
}

/**
 * @param game Game state for the move.
 * @param move The move to execute (all visited cell names).
 * @param online
 */
export async function executeMove(
	game: Game,
	move: string[],
	online: boolean,
): Promise<Game> {
	return zGame.parse(
		await fetchApi(`/api/v1/games/${online ? `${game.uuid}/` : ""}move`, {
			method: "POST",
			body: JSON.stringify({
				game: game,
				move: move,
			}),
		}),
	);
}

export async function newGame(playerName: string): Promise<Game> {
	return zGame.parse(
		await fetchApi("/api/v1/games/new", {
			method: "POST",
			body: JSON.stringify({
				playerName: playerName,
			}),
		}),
	);
}

export async function joinGame(
	gameCode: string,
	playerName: string,
): Promise<Game> {
	return zGame.parse(
		await fetchApi("/api/v1/games/join", {
			method: "POST",
			body: JSON.stringify({
				gameCode: gameCode,
				playerName: playerName,
			}),
		}),
	);
}

export async function evaluateGame(game: Game): Promise<GameEvaluation> {
	return zGameEvaluation.parse(
		await fetchApi("/api/v1/games/evaluate", {
			method: "POST",
			body: JSON.stringify(game),
		}),
	);
}

/**
 * @param game The game for which to find the best move.
 * @param time Maximum time in seconds to find the best move.
 */
export async function getHint(
	game: Game,
	time?: number,
): Promise<CellIdentifier[]> {
	return z.array(zCellIdentifier).parse(
		await fetchApi("/api/v1/games/hint" + (time ? `?time=${time}` : ""), {
			method: "POST",
			body: JSON.stringify(game),
		}),
	);
}

export async function getValidMoves(
	game: Game,
	from: CellIdentifier,
): Promise<PathsTree> {
	return zPathsTree.parse(
		await fetchApi(
			"/api/v1/games/valid-moves?from=" +
				encodeURI(getCellName(from.row, from.column)),
			{
				method: "POST",
				body: JSON.stringify(game),
			},
		),
	);
}

export function useFetchOngoingGames() {
	return useSuspenseQuery({
		queryKey: ["ongoingGames"],
		queryFn: getOngoingGames,
	});
}

/**
 * @param uuid UUID of the game to get.
 */
export function useFetchGame(uuid: string) {
	return useSuspenseQuery({
		queryKey: ["game", uuid],
		queryFn: () => getGame(uuid),
		retry: false,
	});
}

export function useFetchGameEvaluation(game: Game) {
	return useSuspenseQuery({
		queryKey: ["gameEvaluation", game],
		queryFn: async () => !game.winner && (await evaluateGame(game)),
		retry: false,
	});
}
