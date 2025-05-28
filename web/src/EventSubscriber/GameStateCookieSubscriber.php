<?php

namespace App\EventSubscriber;

use App\Game\GameSession;
use App\Game\GameState;
use App\Game\OnlineGame;
use Symfony\Component\EventDispatcher\EventSubscriberInterface;
use Symfony\Component\HttpFoundation\Exception\SessionNotFoundException;
use Symfony\Component\HttpFoundation\RequestStack;
use Symfony\Component\HttpKernel\Event\ResponseEvent;

/**
 * Set the updated game state cookie when sending the response.
 */
readonly class GameStateCookieSubscriber implements EventSubscriberInterface
{
	/**
	 * @param GameState $gameState Game state service.
	 * @param GameSession $gameSession Game session service.
	 * @param OnlineGame $onlineGame Online game service.
	 */
	public function __construct(private GameState $gameState, private GameSession $gameSession, private OnlineGame $onlineGame)
	{
	}

	public function onResponseEvent(ResponseEvent $event): void
	{
		try
		{
			if (!empty($updatedGameState = $this->gameSession->getUpdatedGameState()))
			{ // Create a cookie from the updated game state, if there is one.
				$event->getResponse()->headers->setCookie($this->gameState->createCookie($updatedGameState));
				$this->gameSession->clearUpdatedGameState();
			}
			// Update online games cookie.
			$event->getResponse()->headers->setCookie($this->onlineGame->createOnlineGamesCookie());
		}
		catch (SessionNotFoundException $exception) {}
	}

	public static function getSubscribedEvents(): array
	{
		return [
			ResponseEvent::class => "onResponseEvent",
		];
	}
}
