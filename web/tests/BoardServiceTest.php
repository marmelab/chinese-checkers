<?php

namespace App\Tests;

use App\Game\BoardService;
use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;

class BoardServiceTest extends KernelTestCase
{
	/**
	 * Board utilities.
	 * @var BoardService
	 */
	private BoardService $boardService;

	protected function setUp(): void
	{
		parent::setUp();

		// Boot the Symfony kernel.
		self::bootKernel();

		// Get the board service.
		$this->boardService = static::getContainer()->get(BoardService::class);
	}

	public function testRowName(): void
	{
		$this->assertEquals("a", $this->boardService->getRowName(0), "row name should be \"a\"");
		$this->assertEquals("e", $this->boardService->getRowName(4), "row name should be \"e\"");
		$this->assertEquals("g", $this->boardService->getRowName(6), "row name should be \"g\"");
	}

	public function testGreenTargetArea(): void
	{
		$this->assertTrue($this->boardService->inGreenTargetArea(6, 6));
		$this->assertTrue($this->boardService->inGreenTargetArea(3, 6));
		$this->assertTrue($this->boardService->inGreenTargetArea(6, 3));
		$this->assertTrue($this->boardService->inGreenTargetArea(5, 4));
		$this->assertTrue($this->boardService->inGreenTargetArea(4, 5));

		$this->assertFalse($this->boardService->inGreenTargetArea(0, 0));
		$this->assertFalse($this->boardService->inGreenTargetArea(3, 0));
		$this->assertFalse($this->boardService->inGreenTargetArea(0, 3));
		$this->assertFalse($this->boardService->inGreenTargetArea(1, 2));
		$this->assertFalse($this->boardService->inGreenTargetArea(2, 1));

		$this->assertFalse($this->boardService->inGreenTargetArea(4, 4));
	}

	public function testRedTargetArea(): void
	{
		$this->assertTrue($this->boardService->inRedTargetArea(0, 0));
		$this->assertTrue($this->boardService->inRedTargetArea(3, 0));
		$this->assertTrue($this->boardService->inRedTargetArea(0, 3));
		$this->assertTrue($this->boardService->inRedTargetArea(1, 2));
		$this->assertTrue($this->boardService->inRedTargetArea(2, 1));

		$this->assertFalse($this->boardService->inRedTargetArea(6, 6));
		$this->assertFalse($this->boardService->inRedTargetArea(3, 6));
		$this->assertFalse($this->boardService->inRedTargetArea(6, 3));
		$this->assertFalse($this->boardService->inRedTargetArea(5, 4));
		$this->assertFalse($this->boardService->inRedTargetArea(4, 5));

		$this->assertFalse($this->boardService->inRedTargetArea(4, 4));
	}
}
