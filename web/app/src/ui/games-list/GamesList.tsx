import React from "react";
import {useFetchOngoingGames} from "../../api/games";
import {GameCard} from "./GameCard";

/**
 * Games list component.
 */
export function GamesList() {
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
