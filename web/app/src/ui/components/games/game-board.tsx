import React from "react";
import {Game} from "../../../model/game";
import {CellContent, inGreenTargetArea, inRedTargetArea} from "../../../model/cell";
import {Pawn} from "./pawn";
import {classes} from "../../../utils";

/**
 * Game board component.
 */
export function GameBoard({board}: {
	board: Game["board"];
}) {
	return (
		<table className={"game-board"}>
			<tbody>
			{ // Show all rows of the game board.
				board.map((row, rowIndex) => (
					<GameBoardRow key={rowIndex} row={row} rowIndex={rowIndex}/>
				))
			}
			</tbody>
		</table>
	);
}

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

/**
 * Game board cell.
 */
export function GameBoardCell({rowIndex, cellIndex, cell}: {
	rowIndex: number;
	cellIndex: number;
	cell: CellContent;
}) {
	return (
		<td
			className={classes(inGreenTargetArea(rowIndex, cellIndex) && "green-target", inRedTargetArea(rowIndex, cellIndex) && "red-target")}>
			<button type={"button"}>
				{ // Show a pawn if there is one.
					cell != CellContent.Empty && (
						<Pawn pawn={cell}/>
					)
				}
			</button>
		</td>
	);
}
