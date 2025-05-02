<?php

namespace App\Tests\View;

use App\Entity\Board;
use App\Entity\Player;
use App\Game\BoardUtilities;
use App\Game\GameState;
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
		 * Get the board utilities.
		 * @var BoardUtilities $boardUtilities
		 */
		$boardUtilities = static::getContainer()->get(BoardUtilities::class);

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
			$this->assertAnySelectorTextSame("th", $boardUtilities->getRowName($index));

		// Check the 7 column headers title.
		foreach (range(1, 7) as $index)
			$this->assertAnySelectorTextSame("th", "$index");

		// Check the game cookie.
		$this->assertResponseHasCookie(GameState::COOKIE_NAME);
		$this->assertNotNull($rawBoard = json_decode($client->getCookieJar()->get(GameState::COOKIE_NAME)->getValue()), "the game cookie should contain a valid and decodable JSON");
		$this->assertNotNull(Board::initFromRaw($rawBoard), "the game cookie should successfully instantiate a board");
	}

	/**
	 * Test a simple move on the game board.
	 * @return void
	 */
	public function testGameBoardSimpleMove(): void
	{
		// Create test client.
		$client = static::createClient();

		do
		{ // Get the main view again, until green player starts.
			$client->getCookieJar()->clear();
			$crawler = $client->request("GET", "/");
			$rawBoard = json_decode($client->getCookieJar()->get(GameState::COOKIE_NAME)->getValue());;
		}
		while($rawBoard->currentPlayer != 1);

		// Check that we can select a pawn to move, and that there is no pawn on A5.
		$this->assertSelectorTextSame("aside", "Select a pawn to move");
		$this->assertSelectorNotExists("button[value=\"a5\"] .pawn", "shouldn't have a pawn on a5");

		// Click on A4 to start the move.
		$this->assertSelectorExists("button[value=\"a4\"]:not(:disabled)", "there is a clickable button on a4 cell to start the move");
		$client->submit($crawler->filter("button[value=\"a4\"]:not(:disabled)")->form());
		$crawler = $client->followRedirect();

		// Check that A4 has been added to the move list.
		$this->assertSelectorTextSame("ol.moves > li", "a4", "a4 should have been added to the move list");

		// Click on A5 to end the simple move.
		$this->assertSelectorExists("button[value=\"a5\"]:not(:disabled)", "there is a clickable button on a5 cell to end the simple move");
		$client->submit($crawler->filter("button[value=\"a5\"]:not(:disabled)")->form());
		$client->followRedirect();

		// The pawn has moved to A5, we could select a new pawn to move.
		$this->assertSelectorTextSame("aside", "Select a pawn to move");
		$this->assertSelectorNotExists(".flash .error", "should have no error");
		$this->assertSelectorExists("button[value=\"a5\"] .green.pawn", "should have a green pawn on a5");
	}

	/**
	 * Test a jump move on the game board.
	 * @return void
	 */
	public function testGameBoardJumpMove(): void
	{
		// Create test client.
		$client = static::createClient();

		do
		{ // Get the main view again, until green player starts.
			$client->getCookieJar()->clear();
			$crawler = $client->request("GET", "/");
			$rawBoard = json_decode($client->getCookieJar()->get(GameState::COOKIE_NAME)->getValue());;
		}
		while($rawBoard->currentPlayer != 1);

		// Check that we can select a pawn to move, and that there is no pawn on A5.
		$this->assertSelectorTextSame("aside", "Select a pawn to move");
		$this->assertSelectorNotExists("button[value=\"a5\"] .pawn", "shouldn't have a pawn on a5");

		// Click on A3 to start the move.
		$this->assertSelectorExists("button[value=\"a3\"]:not(:disabled)", "there is a clickable button on a3 cell to start the move");
		$client->submit($crawler->filter("button[value=\"a3\"]:not(:disabled)")->form());
		$crawler = $client->followRedirect();

		// Check that A3 has been added to the move list.
		$this->assertSelectorTextSame("ol.moves > li", "a3", "a3 should have been added to the move list");

		// Click on A5 to add it to the move.
		$this->assertSelectorExists("button[value=\"a5\"]:not(:disabled)", "there is a clickable button on a5 cell to add it to the move");
		$client->submit($crawler->filter("button[value=\"a5\"]:not(:disabled)")->form());
		$crawler = $client->followRedirect();

		// A5 has been added to the move list.
		$this->assertSelectorTextSame("ol.moves > li:first-child", "a3", "a3 should be in the move list");
		$this->assertSelectorTextSame("ol.moves > li:nth-child(2)", "a5", "a5 should be in the move list");
		// Still no pawn on A5.
		$this->assertSelectorNotExists("button[value=\"a5\"] .pawn", "shouldn't have a pawn on a5");

		// Click on the End turn button.
		$client->submit($crawler->selectButton("End turn")->form());
		$client->followRedirect();

		// The pawn has moved to A5, we could select a new pawn to move.
		$this->assertSelectorTextSame("aside", "Select a pawn to move");
		$this->assertSelectorNotExists(".flash .error", "should have no error");
		$this->assertSelectorExists("button[value=\"a5\"] .green.pawn", "should have a green pawn on a5");
	}

	/**
	 * Test an invalid move on the game board.
	 * @return void
	 */
	public function testGameBoardInvalidMove(): void
	{
		// Create test client.
		$client = static::createClient();

		do
		{ // Get the main view again, until green player starts.
			$client->getCookieJar()->clear();
			$crawler = $client->request("GET", "/");
			$rawBoard = json_decode($client->getCookieJar()->get(GameState::COOKIE_NAME)->getValue());;
		}
		while($rawBoard->currentPlayer != 1);

		// Check that we can select a pawn to move.
		$this->assertSelectorTextSame("aside", "Select a pawn to move");

		// Click on A3 to start the move.
		$this->assertSelectorExists("button[value=\"a3\"]:not(:disabled)", "there is a clickable button on a3 cell to start the move");
		$client->submit($crawler->filter("button[value=\"a3\"]:not(:disabled)")->form());
		$crawler = $client->followRedirect();

		// Check that A3 has been added to the move list.
		$this->assertSelectorTextSame("ol.moves > li", "a3", "a3 should have been added to the move list");

		// Click on E4 to add it to the move.
		$this->assertSelectorExists("button[value=\"e4\"]:not(:disabled)", "there is a clickable button on e4 cell to add it to the move");
		$client->submit($crawler->filter("button[value=\"e4\"]:not(:disabled)")->form());
		$crawler = $client->followRedirect();

		// E4 has been added to the move list.
		$this->assertSelectorTextSame("ol.moves > li:first-child", "a3", "a3 should be in the move list");
		$this->assertSelectorTextSame("ol.moves > li:nth-child(2)", "e4", "e4 should be in the move list");
		// No pawn on E4.
		$this->assertSelectorNotExists("button[value=\"e4\"] .pawn", "shouldn't have a pawn on e4");

		// Click on the End turn button.
		$client->submit($crawler->selectButton("End turn")->form());
		$client->followRedirect();

		// Still no pawn on E4, an error has been shown.
		$this->assertSelectorTextContains("aside", "Select a pawn to move");
		$this->assertSelectorTextSame(".flash .error", "'e4' cannot be reached from 'a3'");
		$this->assertSelectorNotExists("button[value=\"e4\"] .pawn", "shouldn't have a pawn on e4");
	}
}
