import React from "react";
import { CellIdentifier } from "../../model/cell";
import { MoveOverlay } from "./MoveOverlay";
import { MoveState } from "../board/PlayableGameBoard";
import { clsx } from "clsx";

export function MoveOverlays({
	move,
	hint,
	green,
	red,
}: {
	/**
	 * The move for which to show overlays.
	 */
	move?: MoveState;

	hint?: boolean;

	green?: boolean;
	red?: boolean;
}) {
	hint = !!hint;

	const moveOverlays: React.ReactElement[] = [];
	if (move) {
		// Create move overlays by finding all atomic move pairs (move between a cell and another).
		let previousCell: CellIdentifier = null;
		for (const cell of move) {
			if (previousCell) {
				moveOverlays.push(
					<MoveOverlay
						key={`overlay-${previousCell.row}-${previousCell.column}-${cell.row}-${cell.column}`}
						className={clsx({
							hint: hint,
							green: green,
							red: red,
						})}
						from={previousCell}
						to={cell}
					/>,
				);
			}

			previousCell = cell;
		}
	}

	return moveOverlays;
}
