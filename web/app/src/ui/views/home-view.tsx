import React from "react";
import {usePageTitle} from "../layout";

/**
 * Home view, with ongoing games.
 */
export function HomeView()
{
	usePageTitle("Home");

	return (
		<main>
			<p>Hello World!</p>
		</main>
	);
}
