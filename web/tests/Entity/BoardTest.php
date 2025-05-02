<?php

namespace App\Tests\Entity;

use App\Entity\Board;
use App\Entity\Player;
use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;

class BoardTest extends KernelTestCase
{
	public function testBoardFromRaw(): void
	{
		$rawBoard = (object) [
			"board" => [
				[1, 1, 1, 1, 0, 0, 0],
				[1, 1, 0, 0, 0, 0, 0],
				[1, 1, 0, 0, 0, 0, 2],
				[1, 0, 1, 0, 0, 0, 0],
				[0, 0, 0, 0, 0, 2, 2],
				[0, 0, 0, 0, 2, 2, 2],
				[0, 0, 0, 2, 2, 2, 2],
			],
			"currentPlayer" => 2,
		];

		$board = Board::fromRaw($rawBoard);
		$this->assertEquals($rawBoard->board, $board->getBoard(), "should have the same board");
		$this->assertEquals(Player::from($rawBoard->currentPlayer), $board->getCurrentPlayer(), "should have the same current player");

		$this->assertNull(Board::fromRaw(null), "should be null when raw board is null");
	}
}
