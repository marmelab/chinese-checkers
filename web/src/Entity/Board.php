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
		if (empty($rawBoard)) return null;

		$board = new self();
		$board->board = $rawBoard->board;
		$board->currentPlayer = Player::from($rawBoard->currentPlayer);

		return $board;
	}

	/**
	 * Get the board cells.
	 * @return int[][]
	 */
	public function getBoard(): array
	{
		return $this->board;
	}

	/**
	 * Set the board cells.
	 * @param int[][] $board The new board cells.
	 * @return void
	 */
	public function setBoard(array $board): void
	{
		$this->board = $board;
	}

	/**
	 * Get the current player.
	 * @return Player
	 */
	public function getCurrentPlayer(): Player
	{
		return $this->currentPlayer;
	}

	/**
	 * Set the current player.
	 * @param Player $currentPlayer The new current player.
	 * @return void
	 */
	public function setCurrentPlayer(Player $currentPlayer): void
	{
		$this->currentPlayer = $currentPlayer;
	}
}
