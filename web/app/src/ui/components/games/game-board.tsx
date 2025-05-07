import React from "react";
import {Game} from "../../../model/game";
import {CellContent} from "../../../model/cell";
import {Pawn} from "./pawn";

/**
 * Game board component.
 */
export function GameBoard({board}: {
	board: Game["board"];
})
{
	return (
		<table className={"game-board"}>
			<tbody>
			{ // Show all rows of the game board.
				board.map((row, rowIndex) => (
					<GameBoardRow key={rowIndex} row={row} />
				))
			}
			</tbody>
		</table>
	);
}

/**
 * Game board row of cells.
 */
export function GameBoardRow({row}: {
	row: Game["board"][0];
})
{
	return (
		<tr>
			{ // Show all cells of the row.
				row.map((cell, cellIndex) => (
					<GameBoardCell key={cellIndex} cell={cell} />
				))
			}
		</tr>
	);
}

/**
 * Game board cell.
 */
export function GameBoardCell({cell}: {
	cell: CellContent;
})
{
	return (
		<td>
			<button type={"button"}>
				{ // Show a pawn if there is one.
					cell != CellContent.Empty && (
						<Pawn pawn={cell} />
					)
				}
			</button>
		</td>
	);
}
