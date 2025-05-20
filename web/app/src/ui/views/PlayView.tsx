import React from "react";
import {
	ArrowFatLinesRight,
	GameController,
	Planet,
	Play,
	PlusCircle,
	SignIn,
	UserPlus,
} from "@phosphor-icons/react";
import { Link, useNavigate } from "react-router-dom";
import "./GameView.css";
import { resetLocalGame } from "../../storage/local-game";
import { useAuthenticatedAccount } from "../../storage/authentication";

export function PlayView() {
	const navigate = useNavigate();

	const isAuthenticated = !!useAuthenticatedAccount();

	return (
		<>
			<header>
				<h1>Play</h1>
			</header>
			<main className={"play"}>
				<h2>
					<Planet /> Online
				</h2>

				{isAuthenticated ? (
					<>
						<Link role="button" to={"/app/game/online/new"}>
							<PlusCircle />
							New online game
						</Link>

						<Link role="button" to={"/app/game/online/join"}>
							<ArrowFatLinesRight />
							Join an online game
						</Link>
					</>
				) : (
					<>
						<p className="center">
							You must have an account to play online with your friends.
						</p>
						<Link role="button" to={"/app/account"}>
							<SignIn /> Log in
						</Link>

						<Link role="button" to={"/app/sign-up"}>
							<UserPlus /> Sign up
						</Link>
					</>
				)}

				<h2>
					<GameController /> Local
				</h2>

				<button
					type="button"
					onClick={() => {
						resetLocalGame();
						navigate("/app/game/local");
					}}
				>
					<PlusCircle />
					New local game
				</button>

				<Link role="button" to={"/app/game/local"}>
					<Play />
					Resume local game
				</Link>
			</main>
		</>
	);
}
