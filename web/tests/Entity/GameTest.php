<?php

namespace App\Tests\Entity;

use App\Entity\Game;
use App\Entity\GamePlayer;
use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;

class GameTest extends KernelTestCase
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

		$board = Game::initFromRaw($rawBoard);
		$this->assertEquals($rawBoard->board, $board->getBoard(), "should have the same board");
		$this->assertEquals(GamePlayer::from($rawBoard->currentPlayer), $board->getCurrentPlayer(), "should have the same current player");

		$this->assertNull(Game::initFromRaw(null), "should be null when raw board is null");
	}
}
