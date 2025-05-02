<?php

namespace App\Controller;

use App\Game\BoardUtilities;
use App\Game\GameSession;
use App\Game\GameState;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

/**
 * Game board controller.
 */
class GameController extends AbstractController
{
	/**
	 * The main game board route.
	 * @param GameState $gameState Game state service.
	 * @return Response
	 */
	#[Route("/", name: "game")]
	public function index(GameState $gameState): Response
	{
		// Get the current game state.
		$game = $gameState->getCurrentGame();

		// Initialize a response.
		$response = new Response();
		// Set the updated game cookie.
		$response->headers->setCookie($gameState->createCookie($game));

		// Return the response, with the rendered game.
		return $this->render("game/index.html.twig", [
			"board" => $game->board,
		], $response);
	}
}
