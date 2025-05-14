import React from "react";
import { Link } from "react-router-dom";
import "./GameCard.css";
import "../kit/Card.css";
import { Game, getGameGreenPlayer, getGameRedPlayer } from "../../model/game";
import { Lightning } from "@phosphor-icons/react";

export function GameCard({ game }: { game: Game }) {
	return (
		<Link role="button" className="game" to={`/app/game/${game.uuid}`}>
			<div className="green player">
				{getGameGreenPlayer(game)?.name ?? "Green"}
			</div>
			<div className="vs">
				<Lightning weight="duotone" />
			</div>
			<div className="red player">{getGameRedPlayer(game)?.name ?? "Red"}</div>
		</Link>
	);
}
