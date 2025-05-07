import React from "react";
import {CellContent} from "../../../model/cell";
import {classes} from "../../../utils";

/**
 * Game pawn component.
 */
export function Pawn({pawn}: {
	pawn: CellContent;
})
{
	// Find the class name of the provided pawn.
	const className = pawn == CellContent.GreenPawn ? "green" : (pawn == CellContent.RedPawn ? "red" : undefined);

	return (
		<div className={classes("pawn", className)}></div>
	)
}
