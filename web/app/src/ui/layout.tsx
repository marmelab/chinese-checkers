import React from "react";
import {Menu} from "./menu";
import {Outlet} from "react-router-dom";

/**
 * Layout component.
 */
export function Layout()
{
	return (
		<>
			<header>
				<Menu />
			</header>

			<Outlet />
		</>
	);
}
