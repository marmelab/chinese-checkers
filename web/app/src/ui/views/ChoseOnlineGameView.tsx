import React from "react";
import { Link } from "react-router-dom";
import { ArrowFatLinesRight, PlusCircle } from "@phosphor-icons/react";
import { AuthenticationRequired } from "../accounts/AuthenticationRequired";

export function ChoseOnlineGameView() {
	return (
		<AuthenticationRequired>
			<header>
				<h1>Online Game</h1>
			</header>

			<main className="online game">
				<Link role="button" to={"/app/game/online/new"}>
					<PlusCircle />
					New online game
				</Link>

				<Link role="button" to={"/app/game/online/join"}>
					<ArrowFatLinesRight />
					Join an online game
				</Link>
			</main>
		</AuthenticationRequired>
	);
}
