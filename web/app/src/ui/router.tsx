import {createBrowserRouter} from "react-router-dom";
import {Layout} from "./Layout";
import {HomeView} from "./views/HomeView";
import {GameView} from "./views/GameView";
import {ErrorView} from "./views/ErrorView";

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
				path: "game/:uuid",
				Component: GameView,
			},
		],
	},
]);
