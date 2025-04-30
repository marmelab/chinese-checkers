<?php

namespace App\Twig;

use App\Game\BoardService;
use Twig\Extension\AbstractExtension;
use Twig\TwigFunction;

/**
 * Simple Twig app extension, with useful filters.
 */
class AppExtension extends AbstractExtension
{
	/**
	 * @param BoardService $boardService The board utilities.
	 */
	public function __construct(private BoardService $boardService)
	{}

	public function getFunctions(): array
	{
		return [
			// Function to get a row name from a row index.
			new TwigFunction("rowName", [$this->boardService, "rowName"]),

			// Functions to find out if a provided cell is in a target area.
			new TwigFunction("inGreenTargetArea", [$this->boardService, "inGreenTargetArea"]),
			new TwigFunction("inRedTargetArea", [$this->boardService, "inRedTargetArea"]),
		];
	}
}
