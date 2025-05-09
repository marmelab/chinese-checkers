import {Game} from "../../model/game";
import React from "react";
import {GameBoardCell} from "./GameBoardCell";

/**
 * Game board row of cells.
 */
export function GameBoardRow({rowIndex, row}: {
	rowIndex: number;
	row: Game["board"][0];
}) {
	return (
		<tr>
			{ // Show all cells of the row.
				row.map((cell, cellIndex) => (
					<GameBoardCell key={cellIndex} cell={cell} rowIndex={rowIndex} cellIndex={cellIndex}/>
				))
			}
		</tr>
	);
}
