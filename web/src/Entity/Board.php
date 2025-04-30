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

	/**
	 * Initialize a game board instance with the raw game board data.
	 * @param object|null $rawBoard A raw game board.
	 * @return Board|null The deserialized game board.
	 */
	public static function fromRaw(object|null $rawBoard): Board|null
	{
		$board = new Board();
		$board->board = $rawBoard->board;
		$board->currentPlayer = Player::from($rawBoard->currentPlayer);

		return $board;
	}
}
