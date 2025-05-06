<?php

namespace App\Entity;

use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
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
	#[ORM\GeneratedValue("CUSTOM")]
	#[ORM\CustomIdGenerator("doctrine.uuid_generator")]
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
	 * @var GamePlayer
	 */
	#[ORM\Column(type: "smallint", enumType: GamePlayer::class)]
	private GamePlayer $currentPlayer;

	/**
	 * Related online players.
	 * @var Collection<int, OnlinePlayer>
	 */
	#[ORM\OneToMany(targetEntity: OnlinePlayer::class, mappedBy: "game")]
	#[ORM\JoinColumn(referencedColumnName: "uuid")]
	private Collection $players;

	public function __construct()
	{
		$this->players = new ArrayCollection();
	}

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
	 * @return GamePlayer
	 */
	public function getCurrentPlayer(): GamePlayer
	{
		return $this->currentPlayer;
	}

	/**
	 * Set the current player.
	 * @param GamePlayer $currentPlayer The new current player.
	 * @return void
	 */
	public function setCurrentPlayer(GamePlayer $currentPlayer): void
	{
		$this->currentPlayer = $currentPlayer;
	}

	/**
	 * Get related players of the game.
	 * @return Collection<int, OnlinePlayer>
	 */
	public function getPlayers(): Collection
	{
		return $this->players;
	}

	/**
	 * Find a game player by its UUID.
	 * @param string $uuid The online player UUID.
	 * @return GamePlayer|null Found game player, NULL if it's not a player.
	 */
	public function findGamePlayerByUuid(string $uuid): GamePlayer|null
	{
		return $this->getPlayers()->findFirst(fn (int $_, OnlinePlayer $player) => $player->getUuid() == $uuid)?->getGamePlayer();
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
		$board->currentPlayer = GamePlayer::from($rawBoard->currentPlayer);

		return $board;
	}

	public function jsonSerialize(): mixed
	{
		return get_object_vars($this);
	}
}
