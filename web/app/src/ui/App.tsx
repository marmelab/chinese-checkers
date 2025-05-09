import React from "react";
import {RouterProvider} from "react-router-dom";
import {router} from "./router";
import {IconContext} from "@phosphor-icons/react";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";

/**
 * Application global query client instance.
 */
const appQueryClient = new QueryClient();

/**
 * Main application component.
 */
export function App() {
	return (
		<QueryClientProvider client={appQueryClient}>
			<IconContext value={{
				weight: "bold",
				className: "icon",
				size: "1em",
			}}>
				<main className={"app"}>
					<RouterProvider router={router}/>
				</main>
			</IconContext>
		</QueryClientProvider>
	)
}
