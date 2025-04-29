<?php

namespace App;

/**
 * Default game board.
 */
const DEFAULT_GAME_BOARD = [
	[1, 1, 1, 1, 0, 0, 0],
	[1, 1, 1, 0, 0, 0, 0],
	[1, 1, 0, 0, 0, 0, 0],
	[1, 0, 0, 0, 0, 0, 2],
	[0, 0, 0, 0, 0, 2, 2],
	[0, 0, 0, 0, 2, 2, 2],
	[0, 0, 0, 2, 2, 2, 2],
];

/**
 * Shape of the target area at each side of the board.
 */
const TARGET_AREA_SHAPE = [
	[1, 1, 1, 1],
	[1, 1, 1, 0],
	[1, 1, 0, 0],
	[1, 0, 0, 0],
];
