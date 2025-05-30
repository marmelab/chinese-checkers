import React, { Suspense } from "react";
import "./PlayersScores.css";
import { Game, getGameGreenPlayer, getGameRedPlayer } from "../../model/game";
import { useFetchGameEvaluation } from "../../api/games";
import { Loader } from "../kit/Loader";

export function PlayersScores(props: { game: Game }) {
	return (
		<Suspense fallback={<Loader />}>
			<AsyncPlayersScores {...props} />
		</Suspense>
	);
}

export function AsyncPlayersScores({ game }: { game: Game }) {
	const gameEvaluation = useFetchGameEvaluation(game);

	return (
		<div className="scores">
			<div className="green">
				<strong>{getGameGreenPlayer(game).name}</strong>
				<span className="score">{gameEvaluation.data.evaluation.green}</span>
			</div>
			<div className="red">
				<strong>{getGameRedPlayer(game).name}</strong>
				<span className="score">{gameEvaluation.data.evaluation.red}</span>
			</div>

			<progress
				value={gameEvaluation.data.evaluation.green}
				max={100}
			></progress>
		</div>
	);
}
