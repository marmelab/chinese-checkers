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
import {
	useIsBestMoveHintEnd,
	useIsBestMoveHintStart,
	useIsInBestMoveHint,
} from "../../storage/moves-hint";

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
		move?.[0]?.row == rowIndex && move?.[0]?.column == cellIndex;

	const isPartOfTheMove = !!move?.find(
		(cell) => cell.row == rowIndex && cell.column == cellIndex,
	);

	const isBestMoveHintStart = useIsBestMoveHintStart({
		row: rowIndex,
		column: cellIndex,
	});
	const isInBestMoveHint = useIsInBestMoveHint({
		row: rowIndex,
		column: cellIndex,
	});
	const isBestMoveHintEnd = useIsBestMoveHintEnd({
		row: rowIndex,
		column: cellIndex,
	});

	return (
		<td
			id={`cell-${rowIndex}-${cellIndex}`}
			className={clsx({
				"green-target": inGreenTargetArea(rowIndex, cellIndex),
				"red-target": inRedTargetArea(rowIndex, cellIndex),
				hint: isBestMoveHintEnd,
			})}
		>
			<button
				type="button"
				disabled={!onClick}
				onClick={() => onClick?.(rowIndex, cellIndex)}
			>
				{cell != CellContent.Empty && (
					<Pawn
						pawn={cell}
						selected={isMoveStart}
						hint={isBestMoveHintStart && move?.length == 0}
					/>
				)}
			</button>

			{isPartOfTheMove && <X className="remove-move-part icon" />}
			{isInBestMoveHint && !isBestMoveHintStart && !isBestMoveHintEnd && (
				<span className="best-move-part"></span>
			)}
		</td>
	);
}
