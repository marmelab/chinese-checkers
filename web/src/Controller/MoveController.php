<?php

namespace App\Controller;

use App\Entity\Cell;
use App\Exceptions\GameApiException;
use App\Game\GameSession;
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
	 * @return Response
	 */
	#[Route("/move", name: "move")]
	public function move(Request $request): Response
	{
		// Read the cell from form data.
		$cell = new Cell($request->get("cell"));

		// Check cell validity.
		if ($this->validator->validate($cell)->count() == 0)
		{ // The cell is valid, add it to the move.
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
				$this->addFlash("error", "internal error");
			}
		}
		else
		{ // Invalid cell format.
			$this->addFlash("error", "invalid cell format");
		}

		// Redirect to the view.
		return $this->redirect("/");
	}

	/**
	 * End the current turn.
	 * @return Response
	 */
	#[Route("/move/end", name: "move_end")]
	public function end(): Response
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
			$this->addFlash("error", "internal error");
		}

		// Redirect to the view.
		return $this->redirect("/");
	}
}
