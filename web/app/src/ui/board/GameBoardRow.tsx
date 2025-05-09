import {Game} from "../../model/game";
import React from "react";
import {GameBoardCell} from "./GameBoardCell";

export function GameBoardRow({
	rowIndex,
	row,
	onClick,
}: {
	rowIndex: number;
	row: Game["board"][0];

	onClick?: (rowIndex: number, cellIndex: number) => void;
}) {
	return (
		<tr>
			{
				// Show all cells of the row.
				row.map((cell, cellIndex) => (
					<GameBoardCell
						key={cellIndex}
						cell={cell}
						rowIndex={rowIndex}
						cellIndex={cellIndex}
						onClick={onClick}
					/>
				))
			}
		</tr>
	);
}
