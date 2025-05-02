<?php

namespace App\Tests\Game;

use App\Game\GameSession;
use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\RequestStack;
use Symfony\Component\HttpFoundation\Session\Session;
use Symfony\Component\HttpFoundation\Session\Storage\MockArraySessionStorage;

class GameSessionTest extends KernelTestCase
{
	/**
	 * Game session.
	 * @var GameSession
	 */
	private GameSession $gameSession;

	protected function setUp(): void
	{
		parent::setUp();

		// Boot the Symfony kernel.
		self::bootKernel();

		// Initialize a game session service with a test request.
		$requestStack = new RequestStack();
		$requestStack->push($request = new Request());
		$request->setSession(new Session(new MockArraySessionStorage()));
		$this->gameSession = new GameSession($requestStack);
	}

	/**
	 * Test initial game session.
	 * @return void
	 */
	public function testInitialSessionState(): void
	{
		$this->assertEquals([], $this->gameSession->getCurrentMove(), "should have an empty move list");
		$this->assertEquals(null, $this->gameSession->getMoveStartCell(), "shouldn't have a start cell");
		$this->assertFalse($this->gameSession->isMoveStartCell(0, 0), "shouldn't be the start cell");
		$this->assertFalse($this->gameSession->isMoveStarted(), "move shouldn't be started");
		$this->assertFalse($this->gameSession->isSimpleMove(), "shouldn't be a simple move");
	}

	/**
	 * Test changing the current move.
	 * @return void
	 */
	public function testSetCurrentMove(): void
	{
		$this->assertEquals([], $this->gameSession->getCurrentMove(), "should have an empty move list");

		$this->gameSession->setCurrentMove(["a5", "a6"]);

		$this->assertEquals(["a5", "a6"], $this->gameSession->getCurrentMove(), "should have a move list with a5 and a6");

		$this->gameSession->appendCellToMove("a7");

		$this->assertEquals(["a5", "a6", "a7"], $this->gameSession->getCurrentMove(), "should have a move list with a5, a6 and a7");

		$this->gameSession->resetCurrentMove();

		$this->assertEquals([], $this->gameSession->getCurrentMove(), "should have an empty move list");
	}

	/**
	 * Test move start cell.
	 * @return void
	 */
	public function testMoveStartCell(): void
	{
		$this->assertFalse($this->gameSession->isMoveStartCell(0, 4), "shouldn't be the start cell");
		$this->assertFalse($this->gameSession->isMoveStarted(), "move shouldn't be started");
		$this->assertNull($this->gameSession->getMoveStartCell(), "shouldn't have a start cell");

		$this->gameSession->setCurrentMove(["a5"]);

		$this->assertFalse($this->gameSession->isMoveStartCell(0, 0), "shouldn't be the start cell");
		$this->assertFalse($this->gameSession->isMoveStartCell(5, 2), "shouldn't be the start cell");
		$this->assertTrue($this->gameSession->isMoveStartCell(0, 4), "should be the start cell");
		$this->assertTrue($this->gameSession->isMoveStarted(), "move has started");
		$this->assertEquals("a5", $this->gameSession->getMoveStartCell()?->getName(), "a5 should now be the start cell");

		$this->gameSession->appendCellToMove("a6");

		$this->assertFalse($this->gameSession->isMoveStartCell(0, 5), "shouldn't be the start cell");
		$this->assertTrue($this->gameSession->isMoveStartCell(0, 4), "should still be the start cell");
		$this->assertTrue($this->gameSession->isMoveStarted(), "move has still started");
		$this->assertEquals("a5", $this->gameSession->getMoveStartCell()?->getName(), "a5 should still be the start cell");

		$this->gameSession->resetCurrentMove();

		$this->assertFalse($this->gameSession->isMoveStartCell(0, 4), "shouldn't be the start cell");
		$this->assertFalse($this->gameSession->isMoveStarted(), "move shouldn't be started");
		$this->assertNull($this->gameSession->getMoveStartCell(), "shouldn't have a start cell");
	}

	/**
	 * Test simple move detection.
	 * @return void
	 */
	public function testSimpleMove(): void
	{
		$this->assertFalse($this->gameSession->isSimpleMove(), "shouldn't be a simple move");

		$this->gameSession->setCurrentMove(["a5"]);

		$this->assertFalse($this->gameSession->isSimpleMove(), "still shouldn't be a simple move");

		$this->gameSession->appendCellToMove("a6");

		$this->assertTrue($this->gameSession->isSimpleMove(), "should now be a simple move");

		$this->gameSession->appendCellToMove("a7");

		$this->assertFalse($this->gameSession->isSimpleMove(), "shouldn't be a simple move anymore");

		$this->gameSession->setCurrentMove(["a3", "a5"]);

		$this->assertFalse($this->gameSession->isSimpleMove(), "shouldn't be a simple move");

		$this->gameSession->resetCurrentMove();

		$this->assertFalse($this->gameSession->isSimpleMove(), "shouldn't be a simple move");
	}
}
