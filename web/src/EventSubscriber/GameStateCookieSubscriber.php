<?php

namespace App\EventSubscriber;

use App\Game\GameSession;
use App\Game\GameState;
use Symfony\Component\EventDispatcher\EventSubscriberInterface;
use Symfony\Component\HttpKernel\Event\ResponseEvent;

/**
 * Set the updated game state cookie when sending the response.
 */
readonly class GameStateCookieSubscriber implements EventSubscriberInterface
{
	/**
	 * @param GameState $gameState Game state service.
	 * @param GameSession $gameSession Game session service.
	 */
	public function __construct(private GameState $gameState, private GameSession $gameSession)
	{
	}

	public function onResponseEvent(ResponseEvent $event): void
	{
		if (!empty($updatedGameState = $this->gameSession->getUpdatedGameState()))
		{ // Create a cookie from the updated game state, if there is one.
			$event->getResponse()->headers->setCookie($this->gameState->createCookie($updatedGameState));
			$this->gameSession->clearUpdatedGameState();
		}
	}

	public static function getSubscribedEvents(): array
	{
		return [
			ResponseEvent::class => "onResponseEvent",
		];
	}
}
