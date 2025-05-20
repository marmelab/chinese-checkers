<?php

namespace App\Controller;

use App\Accounts\AccountsManager;
use App\Dto\NewAccount;
use Doctrine\DBAL\Exception\UniqueConstraintViolationException;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Component\Serializer\SerializerInterface;
use Symfony\Component\Validator\Validator\ValidatorInterface;

class ApiAccountsController extends AbstractController
{
	#[Route("/api/v1/accounts", methods: "POST", format: "json")]
	public function create(Request $request, SerializerInterface $serializer, ValidatorInterface $validator, AccountsManager $accountsManager): JsonResponse
	{
		$newAccount = $serializer->deserialize($request->getContent(), NewAccount::class, "json");

		if (($errors = $validator->validate($newAccount))->count() > 0)
			return $this->json($errors, Response::HTTP_BAD_REQUEST);

		try
		{
			return $this->json($accountsManager->create($newAccount));
		}
		catch (UniqueConstraintViolationException $exception)
		{
			$uniqueViolationVarName = str_contains($exception->getMessage(), "uniq_identifier_name") ? "name" : "email";
			return $this->json([
				"error" => "someone with the same $uniqueViolationVarName already exists",
			], Response::HTTP_UNAUTHORIZED);
		}
	}
}
