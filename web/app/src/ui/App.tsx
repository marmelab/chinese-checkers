import React from "react";
import { RouterProvider } from "react-router-dom";
import { IconContext } from "@phosphor-icons/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import "./Fonts.css";
import "./Colors.css";
import "./Common.css";
import "./App.css";
import "./Button.css";
import "./Icons.css";
import "./Forms.css";
import "./Toastify.css";
import { router } from "./router";
import { ToastContainer } from "react-toastify";

/**
 * Application global query client instance.
 */
const appQueryClient = new QueryClient();

export function App() {
	return (
		<QueryClientProvider client={appQueryClient}>
			<IconContext
				value={{
					weight: "bold",
					className: "icon",
					size: "1em",
				}}
			>
				<main className="app">
					<RouterProvider router={router} />
				</main>
				<ToastContainer />
			</IconContext>
		</QueryClientProvider>
	);
}
