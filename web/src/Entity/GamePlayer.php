<?php

namespace App\Entity;

/**
 * Game player.
 */
enum GamePlayer: int
{
	case Green = 1;
	case Red = 2;

	/**
	 * Get a random player.
	 * @return GamePlayer
	 */
	public static function random(): GamePlayer
	{
		return GamePlayer::from(rand(1, 2));
	}
}
