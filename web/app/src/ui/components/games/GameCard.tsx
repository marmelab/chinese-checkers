import React from "react";
import {Card} from "../Card";
import {Game} from "../../../model/game";
import {Link} from "react-router-dom";

/**
 * Game card component.
 */
export function GameCard({game}: {
	game: Game;
})
{
	return (
		<Link className={"game card"} to={`/app/game/${game.uuid}`}>
			<div className={"green player"}>{game.greenPlayer}</div>
			<div className={"vs"}>VS</div>
			<div className={"red player"}>{game.redPlayer}</div>
		</Link>
	)
}
