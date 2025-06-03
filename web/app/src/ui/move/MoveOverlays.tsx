import React from "react";
import { CellIdentifier } from "../../model/cell";
import { MoveOverlay } from "./MoveOverlay";
import { MoveState } from "../board/PlayableGameBoard";
import { clsx } from "clsx";
import { GamePlayer } from "../../model/game-player";

export function MoveOverlays({
	move,
	hint,
	player,
}: {
	/**
	 * The move for which to show overlays.
	 */
	move?: MoveState;

	hint?: boolean;

	player?: GamePlayer;
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
							green: player == GamePlayer.Green,
							red: player == GamePlayer.Red,
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
