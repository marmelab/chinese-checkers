import { createBrowserRouter } from "react-router-dom";
import { Layout } from "./Layout";
import { ErrorView } from "./views/ErrorView";
import { HomeView } from "./views/HomeView";
import { OnlineGameView } from "./views/OnlineGameView";
import { LocalGameView } from "./views/LocalGameView";
import { NewOnlineGameView } from "./views/NewOnlineGameView";
import { JoinOnlineGameView } from "./views/JoinOnlineGameView";
import { PlayView } from "./views/PlayView";
import { MyAccountView } from "./views/MyAccountView";
import { SignUpView } from "./views/SignUpView";
import { RegisteredView } from "./views/RegisteredView";
import { BotGameView } from "./views/BotGameView";

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
				path: "play",
				Component: PlayView,
			},
			{
				path: "game/local",
				Component: LocalGameView,
			},
			{
				path: "game/bot",
				Component: BotGameView,
			},
			{
				path: "game/online/new",
				Component: NewOnlineGameView,
			},
			{
				path: "game/online/join",
				Component: JoinOnlineGameView,
			},
			{
				path: "game/:uuid",
				Component: OnlineGameView,
			},
			{
				path: "account",
				Component: MyAccountView,
			},
			{
				path: "sign-up",
				Component: SignUpView,
			},
			{
				path: "registered",
				Component: RegisteredView,
			},
		],
	},
]);
