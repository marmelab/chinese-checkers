import React, { useMemo } from "react";
import { useFetchOngoingGames } from "../../api/games";
import { GameCard } from "./GameCard";
import { getOnlineGamePlayerId } from "../../storage/online-game";
import { Game, hasPlayer } from "../../model/game";

export function GamesList() {
	const fetchedGames = useFetchOngoingGames();

	const [spectatorGames, myGames] = useMemo(() => {
		const spectatorGames: Game[] = [];
		const myGames: Game[] = [];

		for (const game of fetchedGames.data) {
			if (hasPlayer(game, getOnlineGamePlayerId(game.uuid))) myGames.push(game);
			else spectatorGames.push(game);
		}

		return [spectatorGames, myGames];
	}, [fetchedGames.data]);

	return (
		<>
			{fetchedGames.data?.length > 0 ? (
				<>
					<h2>My games</h2>
					{myGames?.map((game) => <GameCard key={game.uuid} game={game} />)}

					<h2>Games to spectate</h2>
					{spectatorGames?.map((game) => (
						<GameCard key={game.uuid} game={game} />
					))}
				</>
			) : (
				<p className="center">No ongoing game.</p>
			)}
		</>
	);
}
