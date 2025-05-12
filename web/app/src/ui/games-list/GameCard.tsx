import React from "react";
import { Link } from "react-router-dom";
import "./GameCard.css";
import "../kit/Card.css";
import { Game, getGameGreenPlayer, getGameRedPlayer } from "../../model/game";

export function GameCard({ game }: { game: Game }) {
	return (
		<Link className="game card" to={`/app/game/${game.uuid}`}>
			<div className="green player">{getGameGreenPlayer(game).name}</div>
			<div className="vs">VS</div>
			<div className="red player">{getGameRedPlayer(game).name}</div>
		</Link>
	);
}
