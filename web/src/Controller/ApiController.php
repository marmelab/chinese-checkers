<?php

namespace App\Controller;

use App\Entity\Game;
use App\Entity\GamePlayer;
use App\Exceptions\GameApiException;
use App\Game\GameApi;
use App\Game\OnlineGame;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Contracts\HttpClient\Exception\ClientExceptionInterface;
use Symfony\Contracts\HttpClient\Exception\RedirectionExceptionInterface;
use Symfony\Contracts\HttpClient\Exception\ServerExceptionInterface;
use Symfony\Contracts\HttpClient\Exception\TransportExceptionInterface;

class ApiController extends AbstractController
{
	public function __construct(private readonly EntityManagerInterface $entityManager)
	{
	}

	/**
	 * Get all ongoing games.
	 * @return Response
	 */
	#[Route("/api/v1/games", methods: "GET", format: "json")]
	public function getOngoingGames(): Response
	{
		return $this->json(
			$this->entityManager->getRepository(Game::class)->findAllOngoingGames(),
			context: [
				"groups" => "game:read",
			],
		);
	}

	/**
	 * Get a game state from its UUID.
	 * @param string $uuid UUID of the game to get.
	 * @return Response
	 */
	#[Route("/api/v1/games/{uuid}", methods: "GET", format: "json")]
	public function getGame(string $uuid): Response
	{
		$game = $this->entityManager->getRepository(Game::class)->findOneBy(["uuid" => $uuid]);

		if (empty($game)) throw $this->createNotFoundException();

		return $this->json($game, context: [ "groups" => "game:read" ]);
	}

	/**
	 * @param Request $request
	 * @param OnlineGame $onlineGame
	 * @return Response
	 */
	#[Route("/api/v1/games/new", methods: "POST", format: "json")]
	public function newGame(Request $request, OnlineGame $onlineGame): Response
	{
		$body = json_decode($request->getContent());

		if (empty($body?->playerName))
			return $this->json([ "error" => "you must set a player name to create a game" ], 400);

		$game = $onlineGame->newGame($body->playerName);
		return $this->json($game, context: [ "groups" => "game:read" ]);
	}

	/**
	 * @param Request $request
	 * @param OnlineGame $onlineGame
	 * @return Response
	 */
	#[Route("/api/v1/games/join", methods: "POST", format: "json")]
	public function joinGame(Request $request, OnlineGame $onlineGame): Response
	{
		$body = json_decode($request->getContent());

		if (empty($body?->gameCode))
			return $this->json([ "error" => "you must provide the code of the game to join" ], 400);
		if (empty($body?->playerName))
			return $this->json([ "error" => "you must set a player name to join a game" ], 400);

		if (empty($game = $this->entityManager->getRepository(Game::class)->findOneBy([ "joinCode" => strtoupper(trim($body->gameCode)) ])))
			return $this->json([ "error" => "no game for provided code" ], 404);

		if ($game->getPlayers()->count() >= 2)
			return $this->json([ "error" => "the game with ID {$game->getUuid()} is already full, please join another one" ], 400);

		$onlineGame->joinAsPlayer($game, GamePlayer::Red, $body->playerName);

		return $this->json($game, context: [ "groups" => "game:read" ]);
	}

	/**
	 * Execute the provided move for a game state using the game engine and return the updated game state.
	 * @param Request $request The request.
	 * @param GameApi $gameApi Game API service.
	 * @return Response
	 * @throws ClientExceptionInterface
	 * @throws RedirectionExceptionInterface
	 * @throws ServerExceptionInterface
	 * @throws TransportExceptionInterface
	 */
	#[Route("/api/v1/games/move", methods: "POST", format: "json")]
	public function executeMove(Request $request, GameApi $gameApi): Response
	{
		$body = json_decode($request->getContent());
		$game = Game::initFromRaw($body->game);

		try
		{
			$game = $gameApi->move($game, $body->move);
			return $this->json($game, context: [ "groups" => "game:read" ]);
		}
		catch (GameApiException $exception)
		{
			return $this->json([
				"error" => $exception->getMessage(),
			], 400);
		}
	}
}
