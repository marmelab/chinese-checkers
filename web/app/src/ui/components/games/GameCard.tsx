import React from "react";
import {Card} from "../Card";

/**
 * Game card component.
 */
export function GameCard()
{
	const greenPlayer = "Alice";
	const redPlayer = "Bob";

	return (
		<Card className={"game"} role={"button"}>
			<p>
				<span className={"green player"}>{greenPlayer}</span>
				{" "}<span className={"vs"}>VS</span>{" "}
				<span className={"red player"}>{redPlayer}</span>
			</p>
		</Card>
	)
}
