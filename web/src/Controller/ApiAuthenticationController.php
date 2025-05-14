<?php

namespace App\Controller;

use App\Entity\User;
use Lexik\Bundle\JWTAuthenticationBundle\Exception\JWTEncodeFailureException;
use Lexik\Bundle\JWTAuthenticationBundle\Services\JWTManager;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Component\Security\Http\Attribute\CurrentUser;

final class ApiAuthenticationController extends AbstractController
{
	/**
	 * @param User|null $user
	 * @param JWTManager $jwtManager
	 * @return JsonResponse
	 * @throws JWTEncodeFailureException
	 */
	#[Route("/api/v1/authentication", name: "api_authentication", methods: ["POST"])]
	public function index(#[CurrentUser] ?User $user, JWTManager $jwtManager): JsonResponse
	{
		if (empty($user))
		{
			return $this->json([
				"message" => "invalid credentials",
			], Response::HTTP_UNAUTHORIZED);
		}

		return $this->json([
			"token" => $jwtManager->create($user),
		]);
	}

	/**
	 * @param JWTManager $jwtManager
	 * @return JsonResponse
	 * @throws JWTEncodeFailureException
	 */
	#[Route("/api/v1/authentication/refresh", name: "api_authentication_refresh", methods: "GET")]
	public function refresh(JWTManager $jwtManager): JsonResponse
	{
		$this->denyAccessUnlessGranted("ROLE_USER");

		return $this->json([
			"token" => $jwtManager->create($this->getUser()),
		]);
	}
}
