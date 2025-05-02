<?php

namespace App\Game;

use App\Entity\Board;
use App\Exceptions\GameApiException;
use Psr\Container\ContainerExceptionInterface;
use Psr\Container\NotFoundExceptionInterface;
use Symfony\Component\DependencyInjection\Attribute\Autowire;
use Symfony\Component\DependencyInjection\ParameterBag\ContainerBagInterface;
use Symfony\Component\HttpClient\Exception\ClientException;
use Symfony\Contracts\HttpClient\Exception\ClientExceptionInterface;
use Symfony\Contracts\HttpClient\Exception\RedirectionExceptionInterface;
use Symfony\Contracts\HttpClient\Exception\ServerExceptionInterface;
use Symfony\Contracts\HttpClient\Exception\TransportExceptionInterface;
use Symfony\Contracts\HttpClient\HttpClientInterface;

/**
 * Game engine API service.
 */
class GameApi
{
	/**
	 * @param string $baseUrl Base URL of the game engine API.
	 * @param HttpClientInterface $http HTTP client.
	 */
	public function __construct(
		#[Autowire(param: "app.game_engine_api_url")]
		private readonly string $baseUrl,
		private readonly HttpClientInterface $http,
	) {}

	/**
	 * Move a pawn on the provided board.
	 * @param Board $board The board on which to move the pawn.
	 * @param string[] $moveList The list of visited cells in the move path.
	 * @return Board The updated board.
	 * @throws TransportExceptionInterface
	 * @throws ClientExceptionInterface
	 * @throws RedirectionExceptionInterface
	 * @throws ServerExceptionInterface
	 * @throws GameApiException
	 */
	public function move(Board $board, array $moveList): Board
	{
		try
		{
			$response = $this->http->request("POST", $this->baseUrl."/move", [
				// Pass the entire board state.
				"json" => $board,

				"query" => [
					// Format the provided move list in a string like "a1,a2,a3".
					"path" => implode(",", $moveList),
				],
			]);

			// Success, parse the updated board and return it.
			$rawBoard = json_decode($response->getContent());
			return Board::initFromRaw($rawBoard);
		}
		catch (ClientException $exception)
		{ // Error, throw an exception.
			$errorResponse = json_decode($exception->getResponse()->getContent(false));
			throw new GameApiException($errorResponse?->error ?? "unknown error", previous: $exception);
		}
	}
}
