import React from "react";
import {usePageTitle} from "../layout";
import {GameList} from "../components/games/game-list";

/**
 * Home view, with ongoing games.
 */
export function HomeView()
{
	usePageTitle("Home");

	return (
		<main>
			<p className={"center"}>Find a game to spectate!</p>

			<GameList />
		</main>
	);
}
