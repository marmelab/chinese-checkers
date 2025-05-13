import React, { useState } from "react";
import { Check } from "@phosphor-icons/react";
import { useNavigate } from "react-router-dom";
import { newGame } from "../../api/games";

export function NewOnlineGameView() {
	const navigate = useNavigate();

	const [playerName, setPlayerName] = useState("");

	return (
		<>
			<header>
				<h1>New game</h1>
			</header>
			<main className="new-game">
				<p className="center">Tell us your name to create a new game!</p>

				<form
					onSubmit={async (event) => {
						event.preventDefault();
						const game = await newGame(playerName);
						navigate(`/app/game/${game.uuid}`);
					}}
				>
					<label htmlFor="player-name">
						Player name
						<input
							type="text"
							id="player-name"
							name="player-name"
							min={1}
							max={200}
							required={true}
							value={playerName}
							onChange={(event) => setPlayerName(event.currentTarget.value)}
						/>
					</label>

					<button>
						<Check /> Submit
					</button>
				</form>
			</main>
		</>
	);
}
