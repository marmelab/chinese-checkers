import React from "react";
import { NavLink } from "react-router-dom";
import { GameController, House } from "@phosphor-icons/react";
import "./Menu.css";

export function Menu() {
	return (
		<nav className="menu">
			<ul>
				<li>
					<NavLink to="/app" end>
						<House /> Home
					</NavLink>
				</li>
				<li>
					<NavLink to="/app/game/local" end>
						<GameController /> Play
					</NavLink>
				</li>
			</ul>
		</nav>
	);
}
