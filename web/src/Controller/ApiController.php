<?php

namespace App\Controller;

use App\Entity\Game;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Component\Serializer\SerializerInterface;

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
}
