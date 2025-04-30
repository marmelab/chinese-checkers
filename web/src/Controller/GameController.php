<?php

namespace App\Controller;

use App\Game\BoardService;
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
	 * @param BoardService $boardService The board utilities.
	 * @return Response
	 */
	#[Route("/", name: "game")]
	public function index(BoardService $boardService): Response
	{
		return $this->render("game/index.html.twig", [
			"board" => $boardService->getDefaultGameBoard(),
			"targetAreaShape" => $boardService->getTargetAreaShape(),
		]);
	}
}
