<?php

namespace App\Game;

use App\Entity\Board;
use App\Entity\Player;
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
use Symfony\Contracts\HttpClient\ResponseInterface;

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
	 * Call the game engine API.
	 * @param string $endpoint The endpoint to call.
	 * @param Board $board The board to send to the game engine API.
	 * @param array $options Request options.
	 * @return ResponseInterface The API HTTP response.
	 * @throws ClientExceptionInterface
	 * @throws GameApiException
	 * @throws RedirectionExceptionInterface
	 * @throws ServerExceptionInterface
	 * @throws TransportExceptionInterface
	 */
	protected function call(string $endpoint, Board $board, array $options = []): ResponseInterface
	{
		// Do the HTTP request.
		$response = $this->http->request("POST", $this->baseUrl.$endpoint, [
			"json" => $board,
			...$options,
		]);

		if ($response->getStatusCode() >= 200 && $response->getStatusCode() < 300)
		{ // Success, return the response.
			return $response;
		}
		else
		{ // Error, throw an exception.
			$errorResponse = json_decode($response->getContent(false));
			throw new GameApiException($errorResponse?->error ?? "unknown error");
		}
	}

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
		// Call the API and parse the updated board.
		$rawBoard = json_decode($this->call("/move", $board, [
			"query" => [
				// Format the provided move list in a string like "a1,a2,a3".
				"path" => implode(",", $moveList),
			],
		])->getContent());
		return Board::initFromRaw($rawBoard);
	}

	/**
	 * Get the winner of the board, if there is one.
	 * @param Board $board The board for which to check the winner.
	 * @return Player|null The winner, or NULL if there is no winner.
	 * @throws ClientExceptionInterface
	 * @throws GameApiException
	 * @throws RedirectionExceptionInterface
	 * @throws ServerExceptionInterface
	 * @throws TransportExceptionInterface
	 */
	public function getWinner(Board $board): Player|null
	{
		// Call the API and parse the winner.
		$rawPlayer = json_decode($this->call("/winner", $board)->getContent());

		if ($rawPlayer == 0)
			// No winner
			return null;

		// There is a player, return it.
		return Player::from($rawPlayer);
	}
}
