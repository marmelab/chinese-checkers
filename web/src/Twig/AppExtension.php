<?php

namespace App\Twig;

use App\Game\BoardUtilities;
use App\Game\GameSession;
use Twig\Extension\AbstractExtension;
use Twig\TwigFunction;

/**
 * Simple Twig app extension, with useful filters.
 */
class AppExtension extends AbstractExtension
{
	/**
	 * @param BoardUtilities $boardUtilities Board utilities.
	 * @param GameSession $gameSession Game session service.
	 */
	public function __construct(private readonly BoardUtilities $boardUtilities, private readonly GameSession $gameSession)
	{}

	public function getFunctions(): array
	{
		return [
			// Functions to get a row and cell names from a cell indices.
			new TwigFunction("getRowName", [$this->boardUtilities, "getRowName"]),
			new TwigFunction("getCellName", [$this->boardUtilities, "getCellName"]),

			// Functions to find out if a provided cell is in a target area.
			new TwigFunction("inGreenTargetArea", [$this->boardUtilities, "inGreenTargetArea"]),
			new TwigFunction("inRedTargetArea", [$this->boardUtilities, "inRedTargetArea"]),

			// Functions about the current game session.
			new TwigFunction("getCurrentMove", [$this->gameSession, "getMoveList"]),
			new TwigFunction("isMoveStarted", [$this->gameSession, "isMoveStarted"]),
			new TwigFunction("isMoveStartCell", [$this->gameSession, "isMoveStartCell"]),
		];
	}
}
