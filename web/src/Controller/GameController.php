<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use const App\DEFAULT_GAME_BOARD;
use const App\TARGET_AREA_SHAPE;

/**
 * Game board controller.
 */
class GameController extends AbstractController
{
	/**
	 * The main game board route.
	 * @return Response
	 */
	#[Route("/", name: "game")]
	public function index(): Response
	{
		return $this->render("game/index.html.twig", [
			"board" => DEFAULT_GAME_BOARD,
			"targetAreaShape" => TARGET_AREA_SHAPE,
		]);
	}
}
