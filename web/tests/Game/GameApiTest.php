<?php


namespace App\Tests\Game;

use App\Entity\Game;
use App\Entity\GamePlayer;
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
		$this->gameApi = static::getContainer()->get(GameApi::class);
	}

	public function testValidMove(): void
	{
		// Test game board.
		$board = new Game();
		$board->setBoard([
			[1, 1, 1, 1, 0, 0, 0],
			[1, 1, 1, 0, 0, 0, 0],
			[1, 1, 0, 0, 0, 0, 0],
			[1, 0, 0, 0, 0, 0, 2],
			[0, 0, 0, 0, 0, 2, 2],
			[0, 0, 0, 0, 2, 2, 2],
			[0, 0, 0, 2, 2, 2, 2],
		]);
		$board->setCurrentPlayer(GamePlayer::Green);

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
			$this->assertEquals(GamePlayer::Red, $board->getCurrentPlayer(), "the next player to move is red");
		} catch (Throwable $exception)
		{
			$this->fail("unreachable statement");
		}
	}

	public function testInvalidMove(): void
	{
		// Test game board.
		$board = new Game();
		$board->setBoard([
			[1, 1, 1, 1, 0, 0, 0],
			[1, 1, 1, 0, 0, 0, 0],
			[1, 1, 0, 0, 0, 0, 0],
			[1, 0, 0, 0, 0, 0, 2],
			[0, 0, 0, 0, 0, 2, 2],
			[0, 0, 0, 0, 2, 2, 2],
			[0, 0, 0, 2, 2, 2, 2],
		]);
		$board->setCurrentPlayer(GamePlayer::Green);

		try
		{ // Try an invalid move.
			$this->gameApi->move($board, ["a4", "a6"]);
			$this->fail("unreachable statement");
		}
		catch (GameApiException $exception)
		{
			$this->assertEquals("'a6' cannot be reached from 'a4'", $exception->getMessage(), "should catch an exception with the invalid move details");
		}
	}

	public function testNoWinner(): void
	{
		// Test game board.
		$board = new Game();
		$board->setBoard([
			[0, 1, 1, 0, 0, 0, 0],
			[1, 1, 1, 1, 0, 0, 0],
			[1, 1, 0, 0, 0, 0, 0],
			[1, 0, 1, 0, 0, 0, 2],
			[0, 0, 0, 0, 0, 2, 2],
			[0, 0, 2, 2, 2, 2, 0],
			[0, 0, 0, 2, 0, 2, 2],
		]);
		$board->setCurrentPlayer(GamePlayer::Red);

		try
		{ // Try an invalid move.
			$winner = $this->gameApi->getWinner($board);
			$this->assertNull($winner, "should have no winner");
		}
		catch (GameApiException $exception)
		{
			$this->fail("unreachable statement");
		}
	}

	public function testGreenWinner(): void
	{
		// Test game board.
		$board = new Game();
		$board->setBoard([
			[0, 2, 2, 2, 0, 0, 0],
			[2, 2, 2, 2, 0, 0, 0],
			[2, 2, 0, 0, 0, 0, 0],
			[2, 0, 0, 0, 0, 0, 1],
			[0, 0, 0, 0, 0, 1, 1],
			[0, 0, 0, 0, 1, 1, 1],
			[0, 0, 0, 1, 1, 1, 1],
		]);
		$board->setCurrentPlayer(GamePlayer::Red);

		try
		{ // Try an invalid move.
			$winner = $this->gameApi->getWinner($board);
			$this->assertEquals(GamePlayer::Green, $winner, "should have green as a winner");
		}
		catch (GameApiException $exception)
		{
			$this->fail("unreachable statement");
		}
	}
}
