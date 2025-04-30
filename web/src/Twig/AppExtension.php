<?php

namespace App\Twig;

use App\Game\BoardUtilities;
use Twig\Extension\AbstractExtension;
use Twig\TwigFunction;

/**
 * Simple Twig app extension, with useful filters.
 */
class AppExtension extends AbstractExtension
{
	/**
	 * @param BoardUtilities $boardUtilities The board utilities.
	 */
	public function __construct(private BoardUtilities $boardUtilities)
	{}

	public function getFunctions(): array
	{
		return [
			// Function to get a row name from a row index.
			new TwigFunction("getRowName", [$this->boardUtilities, "getRowName"]),

			// Functions to find out if a provided cell is in a target area.
			new TwigFunction("inGreenTargetArea", [$this->boardUtilities, "inGreenTargetArea"]),
			new TwigFunction("inRedTargetArea", [$this->boardUtilities, "inRedTargetArea"]),
		];
	}
}
