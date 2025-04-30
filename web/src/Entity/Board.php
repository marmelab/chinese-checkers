<?php

namespace App\Entity;

/**
 * Game board state.
 */
class Board
{
	/**
	 * The board cells.
	 * @var int[][]
	 */
	public array $board;

	/**
	 * The current player.
	 * @var Player
	 */
	public Player $currentPlayer;
}
