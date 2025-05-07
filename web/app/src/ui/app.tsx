import React from "react";
import {RouterProvider} from "react-router-dom";
import {router} from "./router";
import {IconContext} from "@phosphor-icons/react";

/**
 * Main application component.
 */
export function App()
{
	return (
		<IconContext value={{
			weight: "bold",
			className: "icon",
			size: "1em",
		}}>
			<main className={"app"}>
				<RouterProvider router={router} />
			</main>
		</IconContext>
	)
}
