<?php

namespace App\Controller;

use App\Accounts\AccountsManager;
use App\Entity\Account;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Component\Security\Http\Attribute\CurrentUser;

final class ApiAuthenticationController extends AbstractController
{
	/**
	 * @param Account|null $account
	 * @param AccountsManager $accountsManager
	 * @return JsonResponse
	 */
	#[Route("/api/v1/authentication", name: "api_authentication", methods: ["POST"])]
	public function index(#[CurrentUser] ?Account $account, AccountsManager $accountsManager): JsonResponse
	{
		if (empty($account))
		{
			return $this->json([
				"error" => "Invalid credentials.",
			], Response::HTTP_UNAUTHORIZED);
		}

		$response = new JsonResponse();
		$response->setData([
			"token" => $accountsManager->getAuthenticationToken($this->getUser()),
		]);
		$response->headers->setCookie($accountsManager->getAuthenticationCookie($this->getUser()));
		return $response;
	}

	/**
	 * @param AccountsManager $accountsManager
	 * @return JsonResponse
	 */
	#[Route("/api/v1/authentication/refresh", name: "api_authentication_refresh", methods: "GET")]
	public function refresh(AccountsManager $accountsManager): JsonResponse
	{
		$this->denyAccessUnlessGranted("ROLE_USER");

		$response = new JsonResponse();
		$response->setData([
			"token" => $accountsManager->getAuthenticationToken($this->getUser()),
		]);
		$response->headers->setCookie($accountsManager->getAuthenticationCookie($this->getUser()));
		return $response;
	}
}
