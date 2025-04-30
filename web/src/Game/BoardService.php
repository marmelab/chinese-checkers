<?php

namespace App\Game;

/**
 * Service with game board utilities.
 */
class BoardService
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
	public function rowName(int $rowIndex): string
	{
		// Get the ASCII code of 'a', and shift the ASCII code using the row index.
		$ascii = ord('a') + $rowIndex;
		// Get the character of the ASCII code.
		return chr($ascii);
	}
}
