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
}
