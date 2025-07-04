<?php

namespace App\Entity;

use Doctrine\ORM\Mapping as ORM;
use Symfony\Component\Serializer\Annotation\Groups;
use Symfony\Component\Serializer\Attribute\Ignore;

/**
 * Entity of an online player.
 */
#[ORM\Entity]
class OnlinePlayer
{
	#[ORM\Id]
	#[ORM\GeneratedValue("CUSTOM")]
	#[ORM\CustomIdGenerator("doctrine.uuid_generator")]
	#[ORM\Column]
	#[Groups(["game:read"])]
	private string $uuid;

	/**
	 * Online player name.
	 * @var string
	 */
	#[ORM\Column]
	#[Groups(["game:read"])]
	private string $name;

	/**
	 * Player of the chinese checker game.
	 * @var Game
	 */
	#[ORM\ManyToOne(targetEntity: Game::class, inversedBy: "players")]
	#[ORM\JoinColumn(name: "game_uuid", referencedColumnName: "uuid")]
	#[Ignore]
	private Game $game;

	/**
	 * The player ID (color) in the game.
	 * @var GamePlayer
	 */
	#[ORM\Column]
	#[Groups(["game:read"])]
	private GamePlayer $gamePlayer;

	/**
	 * @var Account|null
	 */
	#[ORM\ManyToOne(Account::class)]
	private ?Account $account = null;

	/**
	 * Get online player UUID.
	 * @return string
	 */
	public function getUuid(): string
	{
		return $this->uuid;
	}

	/**
	 * Get the currently related game.
	 * @return Game
	 */
	public function getGame(): Game
	{
		return $this->game;
	}
	/**
	 * Set the currently related game.
	 * @param Game $game The newly related game.
	 * @return void
	 */
	public function setGame(Game $game): void
	{
		$this->game = $game;
	}

	/**
	 * Get the game player ID (color).
	 * @return GamePlayer
	 */
	public function getGamePlayer(): GamePlayer
	{
		return $this->gamePlayer;
	}
	/**
	 * Set the game player ID (color).
	 * @param GamePlayer $gamePlayer New game player ID (color).
	 * @return void
	 */
	public function setGamePlayer(GamePlayer $gamePlayer): void
	{
		$this->gamePlayer = $gamePlayer;
	}

	/**
	 * Get player name.
	 * @return string
	 */
	public function getName(): string
	{
		return $this->name;
	}

	/**
	 * Set player name.
	 * @param string $name The new player name.
	 * @return void
	 */
	public function setName(string $name): void
	{
		$this->name = $name;
	}

	public function getAccount(): ?Account
	{
		return $this->account;
	}

	public function setAccount(?Account $account): void
	{
		$this->account = $account;
	}
}
