<?php

namespace App\Controller;

use App\Exceptions\GameApiException;
use App\Game\GameApi;
use App\Game\GameState;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
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
	#[Route("/", name: "game")]
	public function index(GameApi $gameApi, GameState $gameState): Response
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
