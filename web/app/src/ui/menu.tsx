import React from "react";
import {NavLink} from "react-router-dom";
import {House} from "@phosphor-icons/react";

/**
 * Main application menu.
 */
export function Menu()
{
	return (
		<nav className={"menu"}>
			<ul>
				<li><NavLink to={"/app"}><House /> Home</NavLink></li>
			</ul>
		</nav>
	);
}
