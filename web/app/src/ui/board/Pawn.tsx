import React from "react";
import {clsx} from "clsx";
import "./Pawn.css";
import {CellContent} from "../../model/cell";

const pawnClassName: Partial<Record<CellContent, string>> = {
	[CellContent.GreenPawn]: "green",
	[CellContent.RedPawn]: "red",
};

export function Pawn({pawn, selected}: {pawn: CellContent; selected: boolean}) {
	return (
		<div
			className={clsx("pawn", selected && "selected", pawnClassName[pawn])}
		></div>
	);
}
