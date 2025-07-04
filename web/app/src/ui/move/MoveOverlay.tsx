import React from "react";
import { CaretRight } from "@phosphor-icons/react";
import "./MoveOverlay.css";
import { CellIdentifier } from "../../model/cell";
import { clsx } from "clsx";

/**
 * Overlay height in px.
 */
const OVERLAY_HEIGHT = 16;

/**
 * An overlay to show a simple move between two cells.
 */
export function MoveOverlay({
	className,
	from,
	to,
}: {
	className?: string;
	from: CellIdentifier;
	to: CellIdentifier;
}) {
	return (
		<div
			className={clsx("move-overlay", className)}
			ref={(element) => {
				if (element) {
					const fromPosition = document
						.getElementById(`cell-${from.row}-${from.column}`)
						?.getBoundingClientRect();
					const toPosition = document
						.getElementById(`cell-${to.row}-${to.column}`)
						?.getBoundingClientRect();

					if (fromPosition && toPosition) {
						const gameBoardPosition = document
							.getElementById("game-board-container")
							.getBoundingClientRect();

						// Center of the origin cell.
						const x1 =
							fromPosition.left -
							gameBoardPosition.left +
							fromPosition.width / 2;
						const y1 =
							fromPosition.top -
							gameBoardPosition.top +
							fromPosition.height / 2;

						// Center of the destination cell.
						const x2 =
							toPosition.left - gameBoardPosition.left + toPosition.width / 2;
						const y2 =
							toPosition.top - gameBoardPosition.top + toPosition.height / 2;

						const distance = Math.sqrt(
							(x2 - x1) * (x2 - x1) + (y2 - y1) * (y2 - y1),
						);

						const angle = Math.atan2(y2 - y1, x2 - x1) * (180 / Math.PI);

						// Center of the overlay.
						const cx = (x1 + x2) / 2 - distance / 2;
						const cy = (y1 + y2) / 2 - OVERLAY_HEIGHT / 2;

						element.style = `top: ${cy}px; left: ${cx}px; width: ${distance}px; transform: rotate(${angle}deg);`;
					}
				}
			}}
		>
			<CaretRight size="0.8em" />
		</div>
	);
}
