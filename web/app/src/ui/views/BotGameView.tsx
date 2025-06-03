import React, { useEffect } from "react";
import "./GameView.css";
import { PlayerTurn } from "../board/PlayerTurn";
import { setLocalGame, useLocalGameStore } from "../../storage/local-game";
import { PlayableGameBoard } from "../board/PlayableGameBoard";
import { PlayersWinChances } from "../board/PlayersWinChances";
import { GetMoveHint } from "../board/GetMoveHint";
import { setBotGame, useBotGameStore } from "../../storage/bot-game";
import { GamePlayer } from "../../model/game-player";
import { Loader } from "../kit/Loader";
import { GameBoard } from "../board/GameBoard";
import { executeMove, getHint } from "../../api/games";
import { Game } from "../../model/game";
import { getCellName } from "../../model/cell";
import { showErrorToast } from "../showErrorToast";

async function moveTheBot(game: Game): Promise<void> {
	try {
		const botMove = await getHint(game);
		const updatedGame = await executeMove(
			game,
			botMove.map((cell) => getCellName(cell.row, cell.column)),
			false,
		);

		setBotGame({
			...game,
			board: updatedGame.board,
			currentPlayer: updatedGame.currentPlayer,
			winner: updatedGame.winner,
			lastMove: updatedGame.lastMove,
		});
	} catch (error) {
		showErrorToast(error);
	}
}

export function BotGameView() {
	const botGame = useBotGameStore();

	const isPlayerTurn = botGame.game.currentPlayer == GamePlayer.Green;

	useEffect(() => {
		if (!isPlayerTurn) {
			moveTheBot(botGame.game);
		}
	}, [isPlayerTurn, botGame.game]);

	return (
		<>
			<header>
				<h1>Robot Game</h1>
			</header>
			<main className={"game"}>
				<PlayersWinChances game={botGame.game} />
				{isPlayerTurn ? (
					<PlayableGameBoard game={botGame.game} onChange={setBotGame} />
				) : (
					<GameBoard board={botGame.game.board} />
				)}
				<PlayerTurn game={botGame.game} />
				{isPlayerTurn ? <GetMoveHint game={botGame.game} /> : <Loader />}
			</main>
		</>
	);
}
