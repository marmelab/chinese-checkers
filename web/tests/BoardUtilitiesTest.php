<?php

namespace App\Tests;

use App\Game\BoardUtilities;
use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;

class BoardUtilitiesTest extends KernelTestCase
{
	/**
	 * Board utilities.
	 * @var BoardUtilities
	 */
	private BoardUtilities $boardUtilities;

	protected function setUp(): void
	{
		parent::setUp();

		// Boot the Symfony kernel.
		self::bootKernel();

		// Get the board service.
		$this->boardUtilities = static::getContainer()->get(BoardUtilities::class);
	}

	public function testRowName(): void
	{
		$this->assertEquals("a", $this->boardUtilities->getRowName(0), "row name should be \"a\"");
		$this->assertEquals("e", $this->boardUtilities->getRowName(4), "row name should be \"e\"");
		$this->assertEquals("g", $this->boardUtilities->getRowName(6), "row name should be \"g\"");
	}

	public function testGreenTargetArea(): void
	{
		$this->assertTrue($this->boardUtilities->inGreenTargetArea(6, 6));
		$this->assertTrue($this->boardUtilities->inGreenTargetArea(3, 6));
		$this->assertTrue($this->boardUtilities->inGreenTargetArea(6, 3));
		$this->assertTrue($this->boardUtilities->inGreenTargetArea(5, 4));
		$this->assertTrue($this->boardUtilities->inGreenTargetArea(4, 5));

		$this->assertFalse($this->boardUtilities->inGreenTargetArea(0, 0));
		$this->assertFalse($this->boardUtilities->inGreenTargetArea(3, 0));
		$this->assertFalse($this->boardUtilities->inGreenTargetArea(0, 3));
		$this->assertFalse($this->boardUtilities->inGreenTargetArea(1, 2));
		$this->assertFalse($this->boardUtilities->inGreenTargetArea(2, 1));

		$this->assertFalse($this->boardUtilities->inGreenTargetArea(4, 4));
	}

	public function testRedTargetArea(): void
	{
		$this->assertTrue($this->boardUtilities->inRedTargetArea(0, 0));
		$this->assertTrue($this->boardUtilities->inRedTargetArea(3, 0));
		$this->assertTrue($this->boardUtilities->inRedTargetArea(0, 3));
		$this->assertTrue($this->boardUtilities->inRedTargetArea(1, 2));
		$this->assertTrue($this->boardUtilities->inRedTargetArea(2, 1));

		$this->assertFalse($this->boardUtilities->inRedTargetArea(6, 6));
		$this->assertFalse($this->boardUtilities->inRedTargetArea(3, 6));
		$this->assertFalse($this->boardUtilities->inRedTargetArea(6, 3));
		$this->assertFalse($this->boardUtilities->inRedTargetArea(5, 4));
		$this->assertFalse($this->boardUtilities->inRedTargetArea(4, 5));

		$this->assertFalse($this->boardUtilities->inRedTargetArea(4, 4));
	}
}
