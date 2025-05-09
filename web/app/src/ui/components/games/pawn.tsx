import React from "react";
import {clsx} from "clsx";
import {CellContent} from "../../../model/cell";

/**
 * Pawn class names map.
 */
const pawnClassName: Partial<Record<CellContent, string>> = {
	[CellContent.GreenPawn]: "green",
	[CellContent.RedPawn]: "red",
};

/**
 * Game pawn component.
 */
export function Pawn({pawn}: {
	pawn: CellContent;
}) {
	return (
		<div className={clsx("pawn", pawnClassName[pawn])}></div>
	)
}
