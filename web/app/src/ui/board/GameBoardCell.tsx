import React from "react";
import clsx from "clsx";
import {
	CellContent,
	inGreenTargetArea,
	inRedTargetArea,
} from "../../model/cell";
import {Pawn} from "./Pawn";

export function GameBoardCell({
	rowIndex,
	cellIndex,
	cell,
}: {
	rowIndex: number;
	cellIndex: number;
	cell: CellContent;
}) {
	return (
		<td
			className={clsx({
				"green-target": inGreenTargetArea(rowIndex, cellIndex),
				"red-target": inRedTargetArea(rowIndex, cellIndex),
			})}
		>
			<button type={"button"}>
				{cell != CellContent.Empty && <Pawn pawn={cell} />}
			</button>
		</td>
	);
}
