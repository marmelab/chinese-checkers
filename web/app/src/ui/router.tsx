import {createBrowserRouter} from "react-router-dom";
import {Layout} from "./layout";
import {HomeView} from "./views/home-view";
import {GameView} from "./views/game-view";
import {ErrorView} from "./views/error-view";

/**
 * The main React router object.
 */
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
