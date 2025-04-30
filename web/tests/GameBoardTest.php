<?php

namespace App\Tests;

use App\Game\BoardService;
use Symfony\Bundle\FrameworkBundle\Test\WebTestCase;

/**
 * Test rendered game board.
 */
class GameBoardTest extends WebTestCase
{
	/**
	 * Test the game board simple view.
	 * @return void
	 */
	public function testGameBoardView(): void
	{
		// Create test client.
		$client = static::createClient();

		/**
		 * Get the board service.
		 * @var BoardService $boardService
		 */
		$boardService = static::getContainer()->get(BoardService::class);

		$client->request("GET", "/");

		$this->assertResponseIsSuccessful();

		// Check game board existence and validity.
		$this->assertSelectorExists("table#game-board.game-board", "should have a game board table with id and class \"game-board\"");
		$this->assertSelectorCount(10, "table#game-board td.red-target > .green.pawn", "should have 10 green pawns on the 10 red target cells in the game board table");
		$this->assertSelectorCount(10, "table#game-board td.green-target > .red.pawn", "should have 10 red pawns on the 10 green target cells in the game board table");
		$this->assertSelectorCount(7*7, "table#game-board td", "there should be 49 board cells in the game board table");
		$this->assertSelectorCount(2*7 + 1, "table#game-board th", "there should be 15 board headers (one empty) in the game board table");

		// Check the 7 row headers.
		foreach (range(0, 6) as $index)
			$this->assertAnySelectorTextSame("th", $boardService->getRowName($index));

		// Check the 7 column headers.
		foreach (range(1, 7) as $index)
			$this->assertAnySelectorTextSame("th", "$index");
	}
}
