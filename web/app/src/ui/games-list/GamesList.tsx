import React from "react";
import {useFetchOngoingGames} from "../../api/games";
import {GameCard} from "./GameCard";

export function GamesList() {
	const fetchedGames = useFetchOngoingGames();

	return (
		<>
			{fetchedGames.data?.map((game) => (
				<GameCard key={game.uuid} game={game} />
			))}
		</>
	);
}
