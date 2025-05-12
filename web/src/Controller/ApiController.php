<?php

namespace App\Controller;

use App\Entity\Game;
use App\Exceptions\GameApiException;
use App\Game\GameApi;
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
			$this->entityManager->getRepository(Game::class)->findAllFullGames(),
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
			]);
		}
	}
}
