<?php

namespace App\Entity;

use Symfony\Component\Validator\Constraints as Assert;

/**
 * A board cell, used in forms.
 */
class Cell
{
	/**
	 * The cell name.
	 * @var string
	 */
	#[Assert\Length(2)]
	protected string $name;

	/**
	 * Create a new cell.
	 * @param string|null $cell The cell value, if there is one.
	 */
	public function __construct(string|null $cell = null)
	{
		if (!empty($cell)) $this->name = $cell;
	}

	/**
	 * Get the cell row index from its name.
	 * @return int The cell row index.
	 */
	public function getRowIndex(): int
	{ // The index is the shift from 'a' ASCII code of the first character.
		return ord($this->name[0]) - ord('a');
	}

	/**
	 * Get the cell column index from its name.
	 * @return int The cell column index.
	 */
	public function getColumnIndex(): int
	{ // The index is the shift from '1' ASCII code of the second character.
		return ord($this->name[1]) - ord('1');
	}

	/**
	 * Get cell name.
	 * @return string The cell name.
	 */
	public function getName(): string
	{
		return $this->name;
	}

	/**
	 * Get the difference count with another cell.
	 * @param Cell $cell The cell to compare.
	 * @return int Cell difference.
	 */
	public function diff(Cell $cell): int
	{
		return abs($this->getColumnIndex() - $cell->getColumnIndex()) + abs($this->getRowIndex() - $cell->getRowIndex());
	}
}
