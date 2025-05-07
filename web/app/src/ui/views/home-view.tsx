import React from "react";
import {usePageTitle} from "../layout";
import {GamesList} from "../components/games/GamesList";

/**
 * Home view, with ongoing games.
 */
export function HomeView()
{
	usePageTitle("Home");

	return (
		<main>
			<p className={"center"}>Find a game to spectate!</p>

			<GamesList />
		</main>
	);
}
