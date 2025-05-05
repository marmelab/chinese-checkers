<?php

namespace App\Game;

use App\Entity\Game;
use Doctrine\ORM\EntityManagerInterface;

/**
 * Online game service.
 */
class OnlineGame
{
	/**
	 * @param GameState $gameState Game state service.
	 * @param EntityManagerInterface $entityManager Entity manager service.
	 */
	public function __construct(private GameState $gameState, private EntityManagerInterface $entityManager)
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

		return $game;
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
}
