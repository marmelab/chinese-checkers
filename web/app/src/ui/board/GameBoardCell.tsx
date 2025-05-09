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
	onClick,
}: {
	rowIndex: number;
	cellIndex: number;
	cell: CellContent;

	onClick?: (rowIndex: number, cellIndex: number) => void;
}) {
	return (
		<td
			id={`cell-${rowIndex}-${cellIndex}`}
			className={clsx({
				"green-target": inGreenTargetArea(rowIndex, cellIndex),
				"red-target": inRedTargetArea(rowIndex, cellIndex),
			})}
		>
			<button
				type={"button"}
				disabled={!onClick}
				onClick={() => onClick?.(rowIndex, cellIndex)}
			>
				{cell != CellContent.Empty && <Pawn pawn={cell} />}
			</button>
		</td>
	);
}
