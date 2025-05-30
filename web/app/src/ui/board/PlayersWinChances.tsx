import React, { Suspense } from "react";
import "./PlayersWinChances.css";
import { Game, getGameGreenPlayer, getGameRedPlayer } from "../../model/game";
import { useFetchGameEvaluation } from "../../api/games";
import { Loader } from "../kit/Loader";

export function PlayersWinChances(props: { game: Game }) {
	return (
		<Suspense fallback={<Loader />}>
			<AsyncPlayersWinChances {...props} />
		</Suspense>
	);
}

export function AsyncPlayersWinChances({ game }: { game: Game }) {
	const gameEvaluation = useFetchGameEvaluation(game);

	return (
		<div className="win-chances">
			<div className="green">
				<strong>{getGameGreenPlayer(game).name}</strong>
				<span className="chance">{gameEvaluation.data.evaluation.green}</span>
			</div>
			<div className="red">
				<strong>{getGameRedPlayer(game).name}</strong>
				<span className="chance">{gameEvaluation.data.evaluation.red}</span>
			</div>

			<progress
				value={gameEvaluation.data.evaluation.green}
				max={100}
			></progress>
		</div>
	);
}
