<?php

namespace App\Controller;

use App\Exceptions\GameApiException;
use App\Game\GameApi;
use App\Game\GameSession;
use App\Game\GameState;
use App\Game\OnlineGame;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Contracts\HttpClient\Exception\ClientExceptionInterface;
use Symfony\Contracts\HttpClient\Exception\RedirectionExceptionInterface;
use Symfony\Contracts\HttpClient\Exception\ServerExceptionInterface;
use Symfony\Contracts\HttpClient\Exception\TransportExceptionInterface;

/**
 * Game board controller.
 */
class GameController extends AbstractController
{
	/**
	 * Create a local or online game.
	 * @param Request $request The request.
	 * @param OnlineGame $onlineGame Online game service.
	 * @param GameSession $gameSession Game session service.
	 * @return Response
	 */
	#[Route("/new", name: "newGame")]
	public function create(Request $request, OnlineGame $onlineGame, GameSession $gameSession): Response
	{
		if ($request->isMethod("POST"))
		{
			if ($request->get("type") === "local")
			{ // If we want to create a local game, reset the game state and redirect to the local game page.
				$gameSession->resetGameState();
				return $this->redirectToRoute("localGame");
			}
			else
			{ // We want to create an online game.
				$newGame = $onlineGame->newGame();
				return $this->redirectToRoute("onlineGame", [
					"gameId" => $newGame->getUuid(),
				]);
			}
		}

		return $this->render("game/new.html.twig");
	}

	/**
	 * The online game board route.
	 * @param string $gameId The game UUID.
	 * @param GameApi $gameApi Game API service.
	 * @param OnlineGame $onlineGame Online game service.
	 * @return Response
	 * @throws ClientExceptionInterface
	 * @throws GameApiException
	 * @throws RedirectionExceptionInterface
	 * @throws ServerExceptionInterface
	 * @throws TransportExceptionInterface
	 */
	#[Route("/game/{gameId}", name: "onlineGame")]
	public function online(string $gameId, GameApi $gameApi, OnlineGame $onlineGame): Response
	{
		// Try to find the game from its UUID.
		if (empty($game = $onlineGame->findGame($gameId)))
			throw $this->createNotFoundException("No game for ID $gameId");

		// Return the rendered game.
		return $this->render("game/index.html.twig", [
			"board" => $game->getBoard(),
			"winner" => $gameApi->getWinner($game)?->value,
		]);
	}

	/**
	 * The main game board route.
	 * @param GameApi $gameApi Game API service.
	 * @param GameState $gameState Game state service.
	 * @return Response
	 * @throws GameApiException
	 * @throws ClientExceptionInterface
	 * @throws RedirectionExceptionInterface
	 * @throws ServerExceptionInterface
	 * @throws TransportExceptionInterface
	 */
	#[Route("/local", name: "localGame")]
	public function local(GameApi $gameApi, GameState $gameState): Response
	{
		$game = $gameState->getCurrentGame();

		// Initialize a response.
		$response = new Response();
		// Set the updated game cookie.
		$response->headers->setCookie($gameState->createCookie($game));

		// Return the response, with the rendered game.
		return $this->render("game/index.html.twig", [
			"board" => $game->getBoard(),
			"winner" => $gameApi->getWinner($game)?->value,
		], $response);
	}
}
