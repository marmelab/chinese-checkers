<?php

namespace App\Tests\Game;

use App\Entity\Player;
use App\Game\BoardUtilities;
use App\Game\GameState;
use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\RequestStack;

class GameStateTest extends KernelTestCase
{
	/**
	 * Game state.
	 * @var GameState
	 */
	private GameState $gameState;

	/**
	 * Test request.
	 * @var Request
	 */
	private Request $request;

	protected function setUp(): void
	{
		parent::setUp();

		// Boot the Symfony kernel.
		self::bootKernel();

		// Initialize a game session service with a test request.
		$requestStack = new RequestStack();
		$requestStack->push($this->request = new Request());
		$this->gameState = new GameState($requestStack, new BoardUtilities());
	}

	/**
	 * Test to get the current game state.
	 * @return void
	 */
	public function testCurrentGame(): void
	{
		$game = $this->gameState->getCurrentGame();

		$this->assertEquals((new BoardUtilities())->getDefaultGameBoard(), $game->getBoard(), "the default current game should use the default game board");
		$this->assertEquals(Player::GREEN, $game->getCurrentPlayer(), "the default current game should have green as starting player");
	}

	public function testGameStateCookie(): void
	{
		// Get a game and move a green pawn.
		$game = $this->gameState->getCurrentGame();
		$board = $game->getBoard();
		$board[0][3] = 0;
		$board[0][4] = 1;
		$game->setBoard($board);
		$game->setCurrentPlayer(Player::RED);

		$cookie = $this->gameState->createCookie($game);

		$this->assertEquals(GameState::COOKIE_NAME, $cookie->getName(), "should have the game state cookie name");
		$rawBoard = json_decode($cookie->getValue());
		$this->assertEquals([
			[1, 1, 1, 0, 1, 0, 0],
			[1, 1, 1, 0, 0, 0, 0],
			[1, 1, 0, 0, 0, 0, 0],
			[1, 0, 0, 0, 0, 0, 2],
			[0, 0, 0, 0, 0, 2, 2],
			[0, 0, 0, 0, 2, 2, 2],
			[0, 0, 0, 2, 2, 2, 2],
		], $rawBoard->board, "should have a default game board with one moved pawn");
		$this->assertEquals(Player::RED->value, $rawBoard->currentPlayer, "should have red as current player");

		// Set the cookie in the request to test its retrieval.
		$this->request->cookies->set(GameState::COOKIE_NAME, $cookie->getValue());
		$game = $this->gameState->getCurrentGame();

		$this->assertEquals([
			[1, 1, 1, 0, 1, 0, 0],
			[1, 1, 1, 0, 0, 0, 0],
			[1, 1, 0, 0, 0, 0, 0],
			[1, 0, 0, 0, 0, 0, 2],
			[0, 0, 0, 0, 0, 2, 2],
			[0, 0, 0, 0, 2, 2, 2],
			[0, 0, 0, 2, 2, 2, 2],
		], $game->getBoard(), "should have a default game board with one moved pawn");
		$this->assertEquals(Player::RED, $game->getCurrentPlayer(), "should have red as current player");
	}
}
