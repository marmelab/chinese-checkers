import React, { useState } from "react";
import { Check } from "@phosphor-icons/react";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import { joinGame } from "../../api/games";
import { ApiError, formatErrorMessage } from "../../api/api";

const INVALID_JOIN_CODE_ERROR = "no game for provided code";

export function JoinOnlineGameView() {
	const navigate = useNavigate();

	const [gameCode, setGameCode] = useState("");
	const [playerName, setPlayerName] = useState("");

	const [gameCodeError, setGameCodeError] = useState<string | null>(null);

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

						try {
							const game = await joinGame(gameCode, playerName);
							navigate(`/app/game/${game.uuid}`);
						} catch (error) {
							if (!(error instanceof ApiError)) throw error;

							const errorMessage = await error.getApiMessage();
							if (errorMessage == INVALID_JOIN_CODE_ERROR) {
								setGameCodeError("Invalid game code.");
							} else {
								toast.error(formatErrorMessage(errorMessage));
							}
						}
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
							onChange={(event) => {
								setGameCode(event.currentTarget.value.toUpperCase());
								setGameCodeError(null);
							}}
						/>
						{gameCodeError && <span className="error">{gameCodeError}</span>}
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
