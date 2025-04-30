<?php

namespace App\Game;

use Symfony\Component\HttpFoundation\RequestStack;

/**
 * Game session service.
 */
class GameSession
{
	/**
	 * The current move attribute name in session store.
	 */
	const string CURRENT_MOVE_ATTRIBUTE_NAME = "currentMove";

	/**
	 * @param RequestStack $requestStack Request stack.
	 */
	public function __construct(private RequestStack $requestStack)
	{
	}

	/**
	 * Get the current move.
	 * Return an empty array if nothing is currently saved.
	 * @return string[] Move path (all visited cells, with origin cell as the first element).
	 */
	public function getCurrentMove(): array
	{
		return $this->requestStack->getCurrentRequest()->getSession()->get(self::CURRENT_MOVE_ATTRIBUTE_NAME, []);
	}

	/**
	 * Set the current move.
	 * @param array $moveList Move path (all visited cells, with origin cell as the first element).
	 * @return void
	 */
	public function setCurrentMove(array $moveList): void
	{
		$this->requestStack->getCurrentRequest()->getSession()->set(self::CURRENT_MOVE_ATTRIBUTE_NAME, $moveList);
	}

	/**
	 * Reset the current move list.
	 * @return void
	 */
	public function resetCurrentMove(): void
	{
		$this->requestStack->getCurrentRequest()->getSession()->remove(self::CURRENT_MOVE_ATTRIBUTE_NAME);
	}
}
