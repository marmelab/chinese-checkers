<?php

namespace App\Twig;

use Twig\Extension\AbstractExtension;
use Twig\TwigFilter;

/**
 * Simple Twig app extension, with useful filters.
 */
class AppExtension extends AbstractExtension
{
	public function getFilters(): array
	{
		return [
			new TwigFilter("ord", fn(string $chr) => ord($chr)),
			new TwigFilter("chr", fn(string $ord) => chr($ord)),
		];
	}
}
