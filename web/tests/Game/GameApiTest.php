<?php


namespace App\Tests\Game;

use App\Entity\Board;
use App\Entity\Player;
use App\Exceptions\GameApiException;
use App\Game\GameApi;
use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;

class GameApiTest extends KernelTestCase
{
	/**
	 * Game API.
	 * @var GameApi
	 */
	private GameApi $gameApi;

	protected function setUp(): void
	{
		parent::setUp();

		// Boot the Symfony kernel.
		self::bootKernel();

		// Get the game API service.
		$this->gameApi = static::getContainer()->get(GameApi::class . ".test");
	}

	public function testValidMove(): void
	{
		// Test game board.
		$board = new Board();
		$board->setBoard([
			[1, 1, 1, 1, 0, 0, 0],
			[1, 1, 1, 0, 0, 0, 0],
			[1, 1, 0, 0, 0, 0, 0],
			[1, 0, 0, 0, 0, 0, 2],
			[0, 0, 0, 0, 0, 2, 2],
			[0, 0, 0, 0, 2, 2, 2],
			[0, 0, 0, 2, 2, 2, 2],
		]);
		$board->setCurrentPlayer(Player::GREEN);

		try
		{ // Try a valid move.
			$board = $this->gameApi->move($board, ["a4", "a5"]);

			$this->assertEquals([
				[1, 1, 1, 0, 1, 0, 0],
				[1, 1, 1, 0, 0, 0, 0],
				[1, 1, 0, 0, 0, 0, 0],
				[1, 0, 0, 0, 0, 0, 2],
				[0, 0, 0, 0, 0, 2, 2],
				[0, 0, 0, 0, 2, 2, 2],
				[0, 0, 0, 2, 2, 2, 2],
			], $board->getBoard(), "a pawn has moved from a4 to a5");
			$this->assertEquals(Player::RED, $board->getCurrentPlayer(), "the next player to move is red");
		} catch (Throwable $exception)
		{
			$this->fail("unreachable statement");
		}
	}

	public function testInvalidMove(): void
	{
		// Test game board.
		$board = new Board();
		$board->setBoard([
			[1, 1, 1, 1, 0, 0, 0],
			[1, 1, 1, 0, 0, 0, 0],
			[1, 1, 0, 0, 0, 0, 0],
			[1, 0, 0, 0, 0, 0, 2],
			[0, 0, 0, 0, 0, 2, 2],
			[0, 0, 0, 0, 2, 2, 2],
			[0, 0, 0, 2, 2, 2, 2],
		]);
		$board->setCurrentPlayer(Player::GREEN);

		try
		{ // Try an invalid move.
			$this->gameApi->move($board, ["a4", "a6"]);
			$this->fail("unreachable statement");
		} catch (GameApiException $exception)
		{
			$this->assertEquals("'a6' cannot be reached from 'a4'", $exception->getMessage(), "should catch an exception with the invalid move details");
		}
	}
}
