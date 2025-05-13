import React from "react";
import { GamesList } from "../games-list/GamesList";

export function HomeView() {
	return (
		<>
			<header>
				<h1>Home</h1>
			</header>
			<main>
				<GamesList />
			</main>
		</>
	);
}
