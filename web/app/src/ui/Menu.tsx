import React from "react";
import {NavLink} from "react-router-dom";
import "./Menu.css";
import {House} from "@phosphor-icons/react";

export function Menu() {
	return (
		<nav className={"menu"}>
			<ul>
				<li>
					<NavLink to={"/app"} end>
						<House /> Home
					</NavLink>
				</li>
			</ul>
		</nav>
	);
}
