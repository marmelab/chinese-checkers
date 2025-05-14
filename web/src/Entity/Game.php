<?php

namespace App\Entity;

use App\Repository\GamesRepository;
use DateTime;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\ORM\Mapping as ORM;
use Symfony\Component\Serializer\Annotation\Groups;

/**
 * Game board state.
 */
#[ORM\Entity(repositoryClass: GamesRepository::class)]
#[ORM\Table("games")]
#[ORM\HasLifecycleCallbacks]
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
	#[Groups(["game:read"])]
	private string $uuid;

	/**
	 * Game creation date and time.
	 * @var DateTime
	 */
	#[ORM\Column(type: "datetimetz")]
	private DateTime $createdAt;

	/**
	 * Game update date and time.
	 * @var DateTime
	 */
	#[ORM\Column(type: "datetimetz")]
	private DateTime $updatedAt;

	/**
	 * Unique join code.
	 * @var string|null
	 */
	#[ORM\Column(type: "string", unique: true, nullable: true)]
	private string|null $joinCode = null;

	/**
	 * The board cells.
	 * @var int[][]
	 */
	#[ORM\Column(type: "json", options: ["jsonb" => true])]
	#[Groups(["game:read"])]
	private array $board;

	/**
	 * The current player.
	 * @var GamePlayer
	 */
	#[ORM\Column(type: "smallint", enumType: GamePlayer::class)]
	#[Groups(["game:read"])]
	private GamePlayer $currentPlayer;

	/**
	 * The winner.
	 * @var GamePlayer|null
	 */
	#[ORM\Column(type: "smallint", nullable: true, enumType: GamePlayer::class)]
	#[Groups(["game:read"])]
	private GamePlayer|null $winner = null;

	/**
	 * Related online players.
	 * @var Collection<int, OnlinePlayer>
	 */
	#[ORM\OneToMany(targetEntity: OnlinePlayer::class, mappedBy: "game", fetch: "EAGER")]
	#[ORM\JoinColumn(referencedColumnName: "uuid")]
	#[Groups(["game:read"])]
	private Collection $players;

	public function __construct()
	{
		$this->players = new ArrayCollection();
	}

	/**
	 * Set the creation date at entity creation.
	 * @return void
	 */
	#[ORM\PrePersist]
	public function onPrePersist(): void
	{
		$this->createdAt = new DateTime("now");
		$this->updatedAt = new DateTime("now");
	}

	/**
	 * Set the update date at entity update.
	 * @return void
	 */
	#[ORM\PreUpdate]
	public function onPreUpdate(): void
	{
		$this->updatedAt = new DateTime("now");
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
	 * Set the unique join code of the game.
	 * @param string $joinCode
	 * @return void
	 */
	public function setJoinCode(string $joinCode): void
	{
		$this->joinCode = $joinCode;
	}

	/**
	 * Get the unique join code of the game.
	 * @return string|null
	 */
	public function getJoinCode(): string|null
	{
		return $this->joinCode;
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
	 * Get the current online player.
	 * @return OnlinePlayer
	 */
	public function getCurrentOnlinePlayer(): OnlinePlayer
	{
		return $this->getPlayers()->findFirst(fn (int $_, OnlinePlayer $onlinePlayer) => (
			$onlinePlayer->getGamePlayer() == $this->getCurrentPlayer()
		));
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

	public function getWinner(): ?GamePlayer
	{
		return $this->winner;
	}

	public function setWinner(?GamePlayer $winner): void
	{
		$this->winner = $winner;
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
