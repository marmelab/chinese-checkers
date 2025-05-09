import React from "react";
import {Game} from "../../model/game";
import {GameBoardRow} from "./GameBoardRow";

export function GameBoard({board}: {
	board: Game["board"];
}) {
	return (
		<table className={"game-board"}>
			<tbody>
			{
				board.map((row, rowIndex) => (
					<GameBoardRow key={rowIndex} row={row} rowIndex={rowIndex}/>
				))
			}
			</tbody>
		</table>
	);
}
