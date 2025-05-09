import React from "react";
import {useQuery} from "@tanstack/react-query";
import {getOngoingGames} from "../../../api/games";
import {GameCard} from "./game-card";
import {Loader} from "../loader";

/**
 * Fetch ongoing games.
 */
function useOngoingGames() {
	return useQuery({queryKey: ["ongoingGames"], queryFn: getOngoingGames});
}

/**
 * Games list component.
 */
export function GameList() {
	const fetchedGames = useOngoingGames();

	if (fetchedGames.isPending)
		return <Loader/>;

	if (fetchedGames.isError)
		throw fetchedGames.error;

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
