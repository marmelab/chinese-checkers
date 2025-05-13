<?php

namespace App\Controller;

use App\Entity\Cell;
use App\Exceptions\GameApiException;
use App\Game\GameSession;
use Psr\Log\LoggerInterface;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Component\Validator\Validator\ValidatorInterface;
use Throwable;

/**
 * Move a pawn on the board.
 */
class MoveController extends AbstractController
{
	/**
	 * @param ValidatorInterface $validator Data validator.
	 * @param GameSession $gameSession Game session service.
	 */
	public function __construct(
		private readonly ValidatorInterface $validator,
		private readonly GameSession        $gameSession,
	) {}

	/**
	 * Continue moving a pawn on the board.
	 * @param Request $request The request.
	 * @param LoggerInterface $logger
	 * @return Response
	 */
	#[Route("/move", name: "move", methods: "POST")]
	public function move(Request $request, LoggerInterface $logger): Response
	{
		// Read the cell from form data, and check its validity.
		$cell = new Cell($request->get("cell"));
		if ($this->validator->validate($cell)->count() > 0)
		{ // Invalid cell format.
			$this->addFlash("error", "invalid cell format");
		}

		// The cell is valid, add it to the move.
		try
		{
			$this->gameSession->addMove($cell);
		}
		catch (GameApiException $apiException)
		{ // Caught an API exception: show the error.
			$this->addFlash("error", $apiException->getMessage());
		}
		catch (Throwable $exception)
		{ // Caught any other exception: show an internal error.
			$logger->error($exception);
			$this->addFlash("error", "internal error");
		}

		// Redirect to the view.
		if ($request->get("gameId") == "local")
			return $this->redirectToRoute("localGame");
		else
			return $this->redirectToRoute("onlineGame", [ "gameId" => $request->get("gameId") ]);
	}

	/**
	 * End the current turn.
	 * @param Request $request The request.
	 * @param LoggerInterface $logger
	 * @return Response
	 */
	#[Route("/move/end", name: "move_end", methods: "POST")]
	public function end(Request $request, LoggerInterface $logger): Response
	{
		try
		{
			$this->gameSession->endTurn();
		}
		catch (GameApiException $apiException)
		{ // Caught an API exception: show the error.
			$this->addFlash("error", $apiException->getMessage());
		}
		catch (Throwable $exception)
		{ // Caught any other exception: show an internal error.
			$logger->error($exception);
			$this->addFlash("error", "internal error");
		}

		// Redirect to the view.
		if ($request->get("gameId") == "local")
			return $this->redirectToRoute("localGame");
		else
			return $this->redirectToRoute("onlineGame", [ "gameId" => $request->get("gameId") ]);
	}
}
