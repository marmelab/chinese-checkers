<?php

namespace App\Tests\Entity;

use App\Entity\Cell;
use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;

class CellTest extends KernelTestCase
{
	public function testCellParse(): void
	{
		$cell = new Cell("b5");
		$this->assertEquals("b5", $cell->getName());
		$this->assertEquals(1, $cell->getRowIndex());
		$this->assertEquals(4, $cell->getColumnIndex());
	}
}
