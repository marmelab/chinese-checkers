<?php

namespace App\Controller;

use App\Entity\Cell;
use App\Exceptions\GameApiException;
use App\Game\GameApi;
use App\Game\GameSession;
use App\Game\GameState;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\RedirectResponse;
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
	 * @param GameState $gameState Game state service.
	 * @param GameSession $gameSession Game session service.
	 * @param GameApi $gameApi Game API service.
	 */
	public function __construct(
		private readonly ValidatorInterface $validator,
		private readonly GameState          $gameState,
		private readonly GameSession        $gameSession,
		private readonly GameApi            $gameApi,
	) {}

	/**
	 * Start moving a pawn on the board.
	 * The sent cell must contain a pawn.
	 * @param Request $request The request.
	 * @return Response
	 */
	#[Route("/move/start", name: "move_start")]
	public function start(Request $request): Response
	{
		// Read the cell from form data.
		$cell = new Cell($request->get("cell"));

		// Check cell validity.
		if ($this->validator->validate($cell)->count() == 0)
		{ // Start a new move.
			$this->gameSession->setCurrentMove([$cell->getName()]);
		}
		else
		{ // Invalid cell format.
			$this->addFlash("error", "invalid cell format");
		}

		// Redirect to the view.
		return $this->redirect("/");
	}

	/**
	 * Continue moving a pawn on the board.
	 * The sent cell must be free.
	 * @param Request $request The request.
	 * @return Response
	 */
	#[Route("/move/continue", name: "move_continue")]
	public function continue(Request $request): Response
	{
		// Initialize the response to send, which redirects to the view.
		$response = new RedirectResponse("/");

		// Read the cell from form data.
		$cell = new Cell($request->get("cell"));

		// Check cell validity.
		if ($this->validator->validate($cell)->count() == 0)
		{ // Continue the move.
			$this->gameSession->appendCellToMove($cell->getName());

			if ($this->gameSession->isSimpleMove())
			{ // If the move is a simple move (to an adjacent cell), end the turn now.
				return $this->endTurn($response);
			}
		}
		else
		{ // Invalid cell format.
			$this->addFlash("error", "invalid cell format");
		}

		return $response;
	}

	/**
	 * End the current turn.
	 * @return Response
	 */
	#[Route("/move/end", name: "move_end")]
	public function end(): Response
	{
		// Initialize the response to send, which redirects to the view.
		$response = new RedirectResponse("/");
		return $this->endTurn($response);
	}

	/**
	 * End the turn with the current move and update the current game state, or add a flash error.
	 * @param RedirectResponse $response The redirect response to alter.
	 * @return RedirectResponse The updated redirect response.
	 */
	public function endTurn(RedirectResponse $response): RedirectResponse
	{
		try
		{
			// Set the game state cookie with the updated game state.
			$response->headers->setCookie(
				$this->gameState->createCookie(
					// Execute the move in the game engine, with the current state and move.
					$this->gameApi->move($this->gameState->getCurrentGame(), $this->gameSession->getCurrentMove())
				)
			);
		}
		catch (GameApiException $apiException)
		{ // Caught an API exception: show the error.
			$this->addFlash("error", $apiException->getMessage());
		}
		catch (Throwable $exception)
		{ // Caught any other exception: show an internal error.
			$this->addFlash("error", "internal error");
		}
		finally
		{ // Reset the move and go back to the view.
			$this->gameSession->resetCurrentMove();
			return $response;
		}
	}
}
