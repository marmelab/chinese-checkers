<?php

namespace App\Game;

use App\Entity\Board;
use App\Entity\Cell;
use App\Exceptions\GameApiException;
use Symfony\Component\HttpFoundation\RequestStack;
use Symfony\Contracts\HttpClient\Exception\ClientExceptionInterface;
use Symfony\Contracts\HttpClient\Exception\RedirectionExceptionInterface;
use Symfony\Contracts\HttpClient\Exception\ServerExceptionInterface;
use Symfony\Contracts\HttpClient\Exception\TransportExceptionInterface;

/**
 * Game session service.
 */
class GameSession
{
	/**
	 * The current move attribute name in session store.
	 */
	const string MOVE_LIST_ATTRIBUTE_NAME = "moveList";

	/**
	 * The updated game state attribute.
	 */
	const string UPDATED_GAME_STATE_ATTRIBUTE_NAME = "updatedGameState";

	/**
	 * @param RequestStack $requestStack Request stack.
	 * @param GameState $gameState Game state service.
	 * @param GameApi $gameApi Game API service.
	 */
	public function __construct(
		private readonly RequestStack $requestStack,
		private readonly GameState    $gameState,
		private readonly GameApi      $gameApi,
	) {}

	/**
	 * Get the current move.
	 * Return an empty array if nothing is currently saved.
	 * @return string[] Move path (all visited cells, with origin cell as the first element).
	 */
	public function getMoveList(): array
	{
		return $this->requestStack->getCurrentRequest()->getSession()->get(self::MOVE_LIST_ATTRIBUTE_NAME, []);
	}

	/**
	 * Get the initial cell of the move.
	 * @return Cell|null The initial cell of the move, NULL if there is no cell in the move.
	 */
	public function getMoveStartCell(): Cell|null
	{
		if (!empty($this->getMoveList()))
			return new Cell($this->getMoveList()[0]);
		else
			return null;
	}

	/**
	 * Determine if the provided cell is the move start.
	 * @param int $rowIndex Row index of the cell.
	 * @param int $columnIndex Column index of the cell.
	 * @return bool True if it is the start of the move list, false otherwise.
	 */
	public function isMoveStartCell(int $rowIndex, int $columnIndex): bool
	{
		// If there is a move start cell...
		return !empty($moveStartCell = $this->getMoveStartCell()) &&
			// ... check that this move start cell indices are the same than the provided ones.
			$moveStartCell->getRowIndex() == $rowIndex && $moveStartCell->getColumnIndex() == $columnIndex;
	}

	/**
	 * Determine if the move is started.
	 * @return bool True if the move is started, false otherwise.
	 */
	public function isMoveStarted(): bool
	{
		return !empty($this->getMoveList());
	}

	/**
	 * Determine if the move is a move to an adjacent cell.
	 * @return bool True if the move is a move to an adjacent cell, false otherwise.
	 */
	public function isSimpleMove(): bool
	{
		// A simple move has only 2 cells in the move list (origin and destination).
		if (count($this->getMoveList()) != 2)
			return false;
		$origin = new Cell($this->getMoveList()[0]);
		$destination = new Cell($this->getMoveList()[1]);

		// A simple move has only one difference (in row or column) between origin and destination.
		return $origin->diff($destination) == 1;
	}

	/**
	 * Set the current move list.
	 * @param array $moveList Move path (all visited cells, with origin cell as the first element).
	 * @return void
	 */
	public function setCurrentMoveList(array $moveList): void
	{
		$this->requestStack->getCurrentRequest()->getSession()->set(self::MOVE_LIST_ATTRIBUTE_NAME, $moveList);
	}

	/**
	 * Append the cell name to the current move list.
	 * @param string $cell Name of the cell to add to the move list.
	 * @return void
	 */
	public function appendCellToMoveList(string $cell): void
	{
		// Get the current move list, and append the provided cell.
		$move = $this->getMoveList();
		$move[] = $cell;
		$this->setCurrentMoveList($move);
	}

	/**
	 * Reset the current move list.
	 * @return void
	 */
	public function resetMoveList(): void
	{
		$this->requestStack->getCurrentRequest()->getSession()->remove(self::MOVE_LIST_ATTRIBUTE_NAME);
	}

	/**
	 * Add a cell to the current move list.
	 * If the move after adding the cell is a simple move (to an adjacent cell), instantly end the turn.
	 * @param Cell $cell The cell to add to the move.
	 * @return void
	 * @throws ClientExceptionInterface
	 * @throws GameApiException
	 * @throws RedirectionExceptionInterface
	 * @throws ServerExceptionInterface
	 * @throws TransportExceptionInterface
	 */
	public function addMove(Cell $cell): void
	{
		$this->appendCellToMoveList($cell->getName());

		if ($this->isSimpleMove())
		{ // If the move is a simple move (to an adjacent cell), end the turn now.
			$this->endTurn();
		}
	}

	/**
	 * End the turn.
	 * Call the game API to check moves validity and move the pawn to the right cell.
	 * @return void
	 * @throws GameApiException
	 * @throws ClientExceptionInterface
	 * @throws RedirectionExceptionInterface
	 * @throws ServerExceptionInterface
	 * @throws TransportExceptionInterface
	 */
	public function endTurn(): void
	{
		try
		{ // Execute the move in the game engine, with the current state and move.
			$updatedGameState = $this->gameApi->move($this->gameState->getCurrentGame(), $this->getMoveList());
			$this->requestStack->getCurrentRequest()->getSession()->set(self::UPDATED_GAME_STATE_ATTRIBUTE_NAME, $updatedGameState);
		}
		finally
		{ // Reset the current move in any case (success or failure).
			$this->resetMoveList();
		}
	}

	/**
	 * Get the updated game state.
	 * @return Board|null Game state.
	 */
	public function getUpdatedGameState(): Board|null
	{
		return $this->requestStack->getCurrentRequest()->getSession()->get(self::UPDATED_GAME_STATE_ATTRIBUTE_NAME);
	}

	/**
	 * Clear the updated game state.
	 * @return void
	 */
	public function clearUpdatedGameState(): void
	{
		$this->requestStack->getCurrentRequest()->getSession()->remove(self::UPDATED_GAME_STATE_ATTRIBUTE_NAME);
	}
}
