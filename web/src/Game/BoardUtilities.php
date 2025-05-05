<?php

namespace App\Game;

/**
 * Service with game board utilities.
 */
class BoardUtilities
{
	/**
	 * Get the default game board.
	 * @return int[][]
	 */
	public function getDefaultGameBoard(): array
	{
		return [
			[1, 1, 1, 1, 0, 0, 0],
			[1, 1, 1, 0, 0, 0, 0],
			[1, 1, 0, 0, 0, 0, 0],
			[1, 0, 0, 0, 0, 0, 2],
			[0, 0, 0, 0, 0, 2, 2],
			[0, 0, 0, 0, 2, 2, 2],
			[0, 0, 0, 2, 2, 2, 2],
		];
	}

	/**
	 * Get the shape of the target area on each side of the 7x7 board.
	 * @return int[][]
	 */
	public function getTargetAreaShape(): array
	{
		return [
			[1, 1, 1, 1],
			[1, 1, 1, 0],
			[1, 1, 0, 0],
			[1, 0, 0, 0],
		];
	}

	/**
	 * Get a row name from its provided index.
	 * Row names are letters starting by 'a'.
	 * @param int $rowIndex The row index.
	 * @return string The row name.
	 */
	public function getRowName(int $rowIndex): string
	{
		// Get the ASCII code of 'a', and shift the ASCII code using the row index.
		$ascii = ord('a') + $rowIndex;
		// Get the character of the ASCII code.
		return chr($ascii);
	}

	/**
	 * Get a cell name from its provided indices.
	 * @param int $rowIndex The row index.
	 * @param int $columnIndex The column index.
	 * @return string The cell name.
	 */
	public function getCellName(int $rowIndex, int $columnIndex): string
	{
		return $this->getRowName($rowIndex).($columnIndex+1);
	}

	/**
	 * Find out if the provided cell (row;column) is in the green target area.
	 * @param int $rowIndex The row index.
	 * @param int $columnIndex The column index.
	 * @return bool True if the provided cell is in the green target area, false otherwise.
	 */
	public function inGreenTargetArea(int $rowIndex, int $columnIndex): bool
	{
		// Get board size.
		$boardSize = count($this->getDefaultGameBoard());
		// The index must exist and be 1 in the reversed target area.
		return !empty($this->getTargetAreaShape()[$boardSize - 1 - $rowIndex]) && !empty($this->getTargetAreaShape()[$boardSize - 1 - $rowIndex][$boardSize - 1 - $columnIndex])
			&& $this->getTargetAreaShape()[$boardSize - 1 - $rowIndex][$boardSize - 1 - $columnIndex];
	}

	/**
	 * Find out if the provided cell (row;column) is in the green target area.
	 * @param int $rowIndex The row index.
	 * @param int $columnIndex The column index.
	 * @return bool True if the provided cell is in the green target area, false otherwise.
	 */
	public function inRedTargetArea(int $rowIndex, int $columnIndex): bool
	{
		// The index must exist and be 1.
		return !empty($this->getTargetAreaShape()[$rowIndex]) && !empty($this->getTargetAreaShape()[$rowIndex][$columnIndex])
			&& $this->getTargetAreaShape()[$rowIndex][$columnIndex];
	}
}
