<?php

namespace App\Game;

use App\Entity\Board;
use App\Entity\Player;
use Symfony\Component\HttpFoundation\Cookie;
use Symfony\Component\HttpFoundation\RequestStack;

/**
 * Game state service.
 */
class GameState
{
	/**
	 * Name of the cookie to use for game state storage.
	 */
	const COOKIE_NAME = "game";

	/**
	 * @param RequestStack $requestStack Request stack.
	 * @param BoardUtilities $boardUtilities Board utilities.
	 */
	public function __construct(private RequestStack $requestStack, private BoardUtilities $boardUtilities)
	{
	}

	/**
	 * Get a new instance of the default game state.
	 * @return Board A default game state.
	 */
	public function getDefaultGame(): Board
	{
		$board = new Board();
		$board->board = $this->boardUtilities->getDefaultGameBoard();
		$board->currentPlayer = Player::GREEN;
		return $board;
	}

	/**
	 * Get the current game state.
	 * If the "game" cookie is set, retrieve the game state from it.
	 * Otherwise, initialize a default game state.
	 * @return Board The game state instance.
	 */
	public function getCurrentGame(): Board
	{
		// Try to get the serialized game from cookies.
		$serializedGame = $this->requestStack->getCurrentRequest()?->cookies?->get(self::COOKIE_NAME);

		if (!empty($serializedGame))
		{ // A serialized game has been found, parse it.
			return Board::fromRaw(json_decode($serializedGame)) ?? $this->getDefaultGame();
		}
		else
		{ // No serialized game.
			return $this->getDefaultGame();
		}
	}

	/**
	 * Create the cookie to store the provided game.
	 * @param Board $game The game to store in cookies.
	 * @return Cookie The cookie to send back in response.
	 */
	public function createCookie(Board $game): Cookie
	{
		return Cookie::create(self::COOKIE_NAME)
			->withValue(json_encode($game))
			->withSecure();
	}
}
