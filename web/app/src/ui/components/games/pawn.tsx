import React from "react";
import {clsx} from "clsx";
import {CellContent} from "../../../model/cell";

/**
 * Game pawn component.
 */
export function Pawn({pawn}: {
	pawn: CellContent;
}) {
	// Find the class name of the provided pawn.
	const className = pawn == CellContent.GreenPawn ? "green" : (pawn == CellContent.RedPawn ? "red" : undefined);

	return (
		<div className={clsx("pawn", className)}></div>
	)
}
