import { createBrowserRouter } from "react-router-dom";
import { Layout } from "./Layout";
import { ErrorView } from "./views/ErrorView";
import { HomeView } from "./views/HomeView";
import { OnlineGameView } from "./views/OnlineGameView";
import { LocalGameView } from "./views/LocalGameView";
import { NewLocalGameView } from "./views/NewLocalGameView";

export const router = createBrowserRouter([
	{
		path: "/app",
		Component: Layout,
		ErrorBoundary: ErrorView,
		children: [
			{
				index: true,
				Component: HomeView,
			},
			{
				path: "game/local/new",
				Component: NewLocalGameView,
			},
			{
				path: "game/local",
				Component: LocalGameView,
			},
			{
				path: "game/:uuid",
				Component: OnlineGameView,
			},
		],
	},
]);
