import React, { useState } from "react";
import { Check } from "@phosphor-icons/react";
import { useNavigate } from "react-router-dom";
import { joinGame } from "../../api/games";

export function JoinOnlineGameView() {
	const navigate = useNavigate();

	const [gameCode, setGameCode] = useState("");
	const [playerName, setPlayerName] = useState("");

	return (
		<>
			<header>
				<h1>Join a Game</h1>
			</header>
			<main className="new-game">
				<p className="center">
					Write the game code to join. A game code is generated when someone
					creates a new game.
				</p>

				<form
					onSubmit={async (event) => {
						event.preventDefault();
						const game = await joinGame(gameCode, playerName);
						navigate(`/app/game/${game.uuid}`);
					}}
				>
					<label htmlFor="game-code">
						Game code
						<input
							type="text"
							id="game-code"
							name="game-code"
							min={1}
							pattern="^[A-Za-z0-9]{6}$"
							required={true}
							value={gameCode}
							onChange={(event) =>
								setGameCode(event.currentTarget.value.toUpperCase())
							}
						/>
					</label>

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
