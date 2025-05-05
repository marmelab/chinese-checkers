<?php

namespace App\Entity;

/**
 * Game player.
 */
enum Player: int
{
	case Green = 1;
	case Red = 2;

	/**
	 * Get a random player.
	 * @return Player
	 */
	public static function random(): Player
	{
		return Player::from(rand(1, 2));
	}
}
