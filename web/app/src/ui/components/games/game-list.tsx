import React from "react";
import {useFetchOngoingGames} from "../../../api/games";
import {GameCard} from "./game-card";

/**
 * Games list component.
 */
export function GameList() {
	const fetchedGames = useFetchOngoingGames();

	return (
		<>
			{ // Show all the fetched games.
				fetchedGames.data?.map(game => (
					<GameCard key={game.uuid} game={game}/>
				))
			}
		</>
	);
}
