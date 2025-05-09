import {z} from "zod";

/**
 * Game player team.
 */
export enum CellContent {
	Empty = 0,
	GreenPawn = 1,
	RedPawn = 2,
}

export const zCellContent = z.nativeEnum(CellContent);

/**
 * Shape of the target area, in each corner of the game board.
 */
const targetAreaShape = [
	[1, 1, 1, 1],
	[1, 1, 1, 0],
	[1, 1, 0, 0],
	[1, 0, 0, 0],
];
const boardSize = 7;

/**
 * Find out if the provided cell (row;column) is in the green target area.
 * @param rowIndex The row index.
 * @param columnIndex The column index.
 */
export function inGreenTargetArea(rowIndex: number, columnIndex: number): boolean {
	return targetAreaShape?.[boardSize - 1 - rowIndex]?.[boardSize - 1 - columnIndex] === 1;
}

/**
 * Find out if the provided cell (row;column) is in the green target area.
 * @param rowIndex The row index.
 * @param columnIndex The column index.
 */
export function inRedTargetArea(rowIndex: number, columnIndex: number): boolean {
	return targetAreaShape?.[rowIndex]?.[columnIndex] === 1;
}
