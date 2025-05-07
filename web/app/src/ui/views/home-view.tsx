import React from "react";
import {usePageTitle} from "../layout";
import {GameCard} from "../components/games/GameCard";

/**
 * Home view, with ongoing games.
 */
export function HomeView()
{
	usePageTitle("Home");

	return (
		<main>
			<p className={"center"}>Find a game to spectate!</p>

			<GameCard />
			<GameCard />
			<GameCard />
		</main>
	);
}
