<?php

namespace App\Game;

use App\Entity\Game;
use App\Entity\GamePlayer;
use App\Entity\OnlinePlayer;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\HttpFoundation\Cookie;
use Symfony\Component\HttpFoundation\RequestStack;

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
	 */
	public function __construct(
		private RequestStack $requestStack,
		private GameState $gameState,
		private EntityManagerInterface $entityManager,
	)
	{
	}

	/**
	 * Initialize a new online game.
	 * @return Game
	 */
	public function newGame(): Game
	{
		$game = $this->gameState->getDefaultGame();

		$this->entityManager->persist($game);
		$this->entityManager->flush();

		// Join the game as the green player.
		$this->joinAsPlayer($game, GamePlayer::Green);

		return $game;
	}

	/**
	 * Create an online player for the provided game.
	 * @param Game $game The online game.
	 * @param GamePlayer $gamePlayer The game player ID of the new player.
	 * @return OnlinePlayer The created online player.
	 */
	public function newOnlinePlayer(Game $game, GamePlayer $gamePlayer): OnlinePlayer
	{
		$onlinePlayer = new OnlinePlayer();
		$onlinePlayer->setGame($game);
		$onlinePlayer->setGamePlayer($gamePlayer);

		$this->entityManager->persist($onlinePlayer);
		$this->entityManager->flush();

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
	 * @return void
	 */
	public function joinAsPlayer(Game $game, GamePlayer $player): void
	{
		$this->registerOnlinePlayer(
			$this->newOnlinePlayer($game, $player)
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
