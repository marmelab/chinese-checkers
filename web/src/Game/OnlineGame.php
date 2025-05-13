<?php

namespace App\Game;

use App\Entity\Game;
use App\Entity\GamePlayer;
use App\Entity\OnlinePlayer;
use Doctrine\ORM\EntityManagerInterface;
use Doctrine\ORM\Exception\ORMException;
use Symfony\Component\HttpFoundation\Cookie;
use Symfony\Component\HttpFoundation\RequestStack;
use Symfony\Component\Mercure\HubInterface;
use Symfony\Component\Mercure\Update;
use Symfony\Component\Serializer\SerializerInterface;

/**
 * Online game service.
 */
class OnlineGame
{
	/**
	 * The online games cookie.
	 */
	const string COOKIE_NAME = "online-games";

	/**
	 * The updated online games attribute name in session.
	 */
	const string UPDATED_ONLINE_GAMES_ATTRIBUTE_NAME = "updatedOnlineGames";

	/**
	 * @param RequestStack $requestStack The request stack service.
	 * @param GameState $gameState Game state service.
	 * @param EntityManagerInterface $entityManager Entity manager service.
	 * @param HubInterface $mercure Mercure hub service.
	 * @param SerializerInterface $serializer Serialization service.
	 */
	public function __construct(
		private RequestStack $requestStack,
		private GameState $gameState,
		private EntityManagerInterface $entityManager,
		private readonly HubInterface           $mercure,
		private readonly SerializerInterface    $serializer,
	)
	{
	}

	/**
	 * Find a unique game join code.
	 * @return string Generated game join code.
	 */
	public function newGameJoinCode(): string
	{
		do
		{ // Generate a new random code until we find a unique one.
			$joinCode = strtoupper(
				substr(bin2hex(openssl_random_pseudo_bytes(16)), 0, 6)
			);

			$existingGame = $this->entityManager->getRepository(Game::class)->findOneBy([
				"joinCode" => $joinCode,
			]);
		}
		while (!empty($existingGame));

		return $joinCode;
	}

	/**
	 * Initialize a new online game.
	 * @param string $playerName The green player name.
	 * @return Game
	 */
	public function newGame(string $playerName): Game
	{
		$game = $this->gameState->getDefaultGame();
		$game->setJoinCode($this->newGameJoinCode());

		$this->entityManager->persist($game);
		$this->entityManager->flush();

		// Join the game as the green player.
		$this->joinAsPlayer($game, GamePlayer::Green, $playerName);

		return $game;
	}

	/**
	 * Create an online player for the provided game.
	 * @param Game $game The online game.
	 * @param GamePlayer $gamePlayer The game player ID of the new player.
	 * @param string $playerName The player name.
	 * @return OnlinePlayer The created online player.
	 * @throws ORMException
	 */
	public function newOnlinePlayer(Game $game, GamePlayer $gamePlayer, string $playerName): OnlinePlayer
	{
		$onlinePlayer = new OnlinePlayer();
		$onlinePlayer->setGame($game);
		$onlinePlayer->setGamePlayer($gamePlayer);
		$onlinePlayer->setName($playerName);

		$this->entityManager->persist($onlinePlayer);
		$this->entityManager->flush();

		$this->entityManager->refresh($game);
		$this->mercure->publish(new Update($game->getUuid(),
			$this->serializer->serialize($game, "json", [ "groups" => ["game:read"] ])
		));

		return $onlinePlayer;
	}

	/**
	 * Find an online game.
	 * @param string $uuid The game UUID.
	 * @return Game|null The online game, or NULL if not found.
	 */
	public function findGame(string $uuid): Game|null
	{
		return $this->entityManager->getRepository(Game::class)->find($uuid);
	}

	/**
	 * Join a game as the provided player.
	 * @param Game $game The game to join.
	 * @param GamePlayer $player The player ID (color) to be.
	 * @param string $playerName The player name.
	 * @return void
	 */
	public function joinAsPlayer(Game $game, GamePlayer $player, string $playerName): void
	{
		$this->registerOnlinePlayer(
			$this->newOnlinePlayer($game, $player, $playerName)
		);
	}

	/**
	 * Register the provided online player in the online games cookie.
	 * @param OnlinePlayer $onlinePlayer The online player to register.
	 * @return void
	 */
	public function registerOnlinePlayer(OnlinePlayer $onlinePlayer): void
	{
		$this->requestStack->getCurrentRequest()->getSession()->set(self::UPDATED_ONLINE_GAMES_ATTRIBUTE_NAME, [
			$onlinePlayer->getGame()->getUuid() => $onlinePlayer->getUuid()
		]);
	}

	/**
	 * Get the player UUID of the current player (using online games cookie).
	 * @param Game $game The game instance.
	 * @return string|null The player UUID in the game, NULL if not a player of the game.
	 */
	public function getPlayerUuid(Game $game): string|null
	{
		return $this->getOnlineGamesCookie()[$game->getUuid()] ?? null;
	}

	/**
	 * Get the online games cookie.
	 * @return array<string, string> Game UUID => Player UUID associative array.
	 */
	private function getOnlineGamesCookie(): array
	{
		return json_decode($this->requestStack->getCurrentRequest()->cookies->get(self::COOKIE_NAME, "[]"), true);
	}

	/**
	 * Create the cookie to store the online games in which the current player is.
	 * @return Cookie The cookie to send back in response.
	 */
	public function createOnlineGamesCookie(): Cookie
	{
		// Create an online games cookie with online games update.
		$cookie = Cookie::create(self::COOKIE_NAME)
			->withValue(json_encode([
				...$this->getOnlineGamesCookie(),
				...$this->getUpdatedOnlineGames(),
			]))
			->withSecure();

		$this->clearUpdatedOnlineGames();

		return $cookie;
	}

	/**
	 * Get the updated game online games.
	 * @return array<string, string> Game UUID => Player UUID associative array.
	 */
	public function getUpdatedOnlineGames(): array
	{
		return $this->requestStack->getCurrentRequest()->getSession()->get(self::UPDATED_ONLINE_GAMES_ATTRIBUTE_NAME, []);
	}

	/**
	 * Clear the updated online games.
	 * @return void
	 */
	public function clearUpdatedOnlineGames(): void
	{
		$this->requestStack->getCurrentRequest()->getSession()->remove(self::UPDATED_ONLINE_GAMES_ATTRIBUTE_NAME);
	}
}
