import React from "react";
import clsx from "clsx";
import {
	CellContent,
	inGreenTargetArea,
	inRedTargetArea,
} from "../../model/cell";
import { Pawn } from "./Pawn";
import { MoveState } from "./PlayableGameBoard";
import { X } from "@phosphor-icons/react";

export function GameBoardCell({
	move,
	rowIndex,
	cellIndex,
	cell,
	onClick,
}: {
	move: MoveState;
	rowIndex: number;
	cellIndex: number;
	cell: CellContent;

	onClick?: (rowIndex: number, cellIndex: number) => void;
}) {
	const isMoveStart =
		move?.[0]?.rowIndex == rowIndex && move?.[0]?.cellIndex == cellIndex;

	const isPartOfTheMove = !!move?.find(
		(cell) => cell.rowIndex == rowIndex && cell.cellIndex == cellIndex,
	);

	return (
		<td
			id={`cell-${rowIndex}-${cellIndex}`}
			className={clsx({
				"green-target": inGreenTargetArea(rowIndex, cellIndex),
				"red-target": inRedTargetArea(rowIndex, cellIndex),
			})}
		>
			<button
				type="button"
				disabled={!onClick}
				onClick={() => onClick?.(rowIndex, cellIndex)}
			>
				{cell != CellContent.Empty && (
					<Pawn pawn={cell} selected={isMoveStart} />
				)}
			</button>

			{isPartOfTheMove && <X className="remove-move-part icon" />}
		</td>
	);
}
