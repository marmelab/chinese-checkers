<?php

namespace App\Entity;

use Doctrine\ORM\Mapping as ORM;

/**
 * Game board state.
 */
#[ORM\Entity]
#[ORM\Table("games")]
class Game implements \JsonSerializable
{
	/**
	 * The game UUID.
	 * @var string
	 */
	#[ORM\Id]
	#[ORM\GeneratedValue]
	#[ORM\Column(type: "uuid")]
	private string $uuid;

	/**
	 * The board cells.
	 * @var int[][]
	 */
	#[ORM\Column(type: "json", options: ["jsonb" => true])]
	private array $board;

	/**
	 * The current player.
	 * @var Player
	 */
	#[ORM\Column(type: "smallint")]
	private Player $currentPlayer;

	/**
	 * Get the game UUID.
	 * @return string
	 */
	public function getUuid(): string
	{
		return $this->uuid;
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

	/**
	 * Initialize a game board instance with the raw game board data.
	 * @param object|null $rawBoard A raw game board.
	 * @return Game|null The deserialized game board.
	 */
	public static function initFromRaw(object|null $rawBoard): Game|null
	{
		if (empty($rawBoard)) return null;

		$board = new self();
		$board->board = $rawBoard->board;
		$board->currentPlayer = Player::from($rawBoard->currentPlayer);

		return $board;
	}

	public function jsonSerialize(): mixed
	{
		return get_object_vars($this);
	}
}
