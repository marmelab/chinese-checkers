<?php

namespace App\Game;

use App\Entity\Cell;
use Symfony\Component\HttpFoundation\RequestStack;

/**
 * Game session service.
 */
class GameSession
{
	/**
	 * The current move attribute name in session store.
	 */
	const string CURRENT_MOVE_ATTRIBUTE_NAME = "currentMove";

	/**
	 * @param RequestStack $requestStack Request stack.
	 */
	public function __construct(private RequestStack $requestStack)
	{
	}

	/**
	 * Get the current move.
	 * Return an empty array if nothing is currently saved.
	 * @return string[] Move path (all visited cells, with origin cell as the first element).
	 */
	public function getCurrentMove(): array
	{
		return $this->requestStack->getCurrentRequest()->getSession()->get(self::CURRENT_MOVE_ATTRIBUTE_NAME, []);
	}

	/**
	 * Get the initial cell of the move.
	 * @return Cell|null The initial cell of the move, NULL if there is no cell in the move.
	 */
	public function getMoveStartCell(): Cell|null
	{
		if (!empty($this->getCurrentMove()))
			return new Cell($this->getCurrentMove()[0]);
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
		return !empty($this->getCurrentMove());
	}

	/**
	 * Determine if the move is a move to an adjacent cell.
	 * @return bool True if the move is a move to an adjacent cell, false otherwise.
	 */
	public function isSimpleMove(): bool
	{
		// A simple move has only 2 cells in the move (origin and destination).
		if (count($this->getCurrentMove()) != 2)
			return false;
		$origin = new Cell($this->getCurrentMove()[0]);
		$destination = new Cell($this->getCurrentMove()[1]);

		// A simple move has only one difference (in row or column) between origin and destination.
		return $origin->diff($destination) == 1;
	}

	/**
	 * Set the current move.
	 * @param array $moveList Move path (all visited cells, with origin cell as the first element).
	 * @return void
	 */
	public function setCurrentMove(array $moveList): void
	{
		$this->requestStack->getCurrentRequest()->getSession()->set(self::CURRENT_MOVE_ATTRIBUTE_NAME, $moveList);
	}

	/**
	 * Append a cell to the current move list.
	 * @param string $cell The cell to add to the list.
	 * @return void
	 */
	public function appendCellToMove(string $cell): void
	{
		// Get the current move, and append the provided cell.
		$move = $this->getCurrentMove();
		$move[] = $cell;
		$this->setCurrentMove($move);
	}

	/**
	 * Reset the current move list.
	 * @return void
	 */
	public function resetCurrentMove(): void
	{
		$this->requestStack->getCurrentRequest()->getSession()->remove(self::CURRENT_MOVE_ATTRIBUTE_NAME);
	}
}
