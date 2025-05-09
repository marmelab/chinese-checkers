import React from "react";
import {usePageTitle} from "../Layout";
import {GamesList} from "../games-list/GamesList";

/**
 * Home view, with ongoing games.
 */
export function HomeView() {
	usePageTitle("Home");

	return (
		<main>
			<p className={"center"}>Find a game to spectate!</p>

			<GamesList/>
		</main>
	);
}
