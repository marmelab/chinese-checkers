<?php

namespace App\Tests;

use App\Entity\Board;
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
		$this->assertSelectorExists(".game-board", "should have a game board");

		// Check board cells existence.
		$this->assertSelectorCount(7*7, ".game-board td", "there should be 49 board cells in the game board table");
		$this->assertSelectorCount(2*7 + 1, ".game-board th", "there should be 15 board headers (one empty) in the game board table");

		// Check pawns existence.
		$this->assertSelectorCount(10, ".green.pawn", "should have 10 green pawns");
		$this->assertSelectorCount(10, ".red.pawn", "should have 10 red pawns");

		// Check target areas existence.
		$this->assertSelectorCount(10, ".green-target", "should have 10 green target cells");
		$this->assertSelectorCount(10, ".red-target", "should have 10 red target cells");

		// Check the 7 row headers title.
		foreach (range(0, 6) as $index)
			$this->assertAnySelectorTextSame("th", $boardService->getRowName($index));

		// Check the 7 column headers title.
		foreach (range(1, 7) as $index)
			$this->assertAnySelectorTextSame("th", "$index");

		// Check the game cookie.
		$this->assertResponseHasCookie("game");
		$this->assertNotNull($rawBoard = json_decode($client->getCookieJar()->get("game")->getValue()), "the game cookie should contain a valid and decodable JSON");
		$this->assertNotNull(Board::fromRaw($rawBoard), "the game cookie should successfully instantiate a board");
	}
}
