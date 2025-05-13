import React from "react";
import { Link, useNavigate } from "react-router-dom";
import { resetLocalGame } from "../../storage/local-game";
import { Play, PlusCircle } from "@phosphor-icons/react";

export function NewLocalGameView() {
	const navigate = useNavigate();

	return (
		<>
			<header>
				<h1>Local Game</h1>
			</header>
			<main className="local game">
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
