import {createBrowserRouter} from "react-router-dom";
import {Layout} from "./layout";
import {HomeView} from "./views/home-view";

/**
 * The main React router object.
 */
export const router = createBrowserRouter([
	{
		path: "/app",
		Component: Layout,
		children: [
			{
				index: true,
				Component: HomeView,
			},
		],
	}
]);
