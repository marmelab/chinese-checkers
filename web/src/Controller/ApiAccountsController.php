<?php

namespace App\Controller;

use App\Accounts\UsersManager;
use App\Dto\NewUser;
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
	public function create(Request $request, SerializerInterface $serializer, ValidatorInterface $validator, UsersManager $usersManager): JsonResponse
	{
		$newUser = $serializer->deserialize($request->getContent(), NewUser::class, "json");

		if (($errors = $validator->validate($newUser))->count() > 0)
			return $this->json($errors, Response::HTTP_BAD_REQUEST);

		try
		{
			return $this->json($usersManager->create($newUser));
		}
		catch (UniqueConstraintViolationException $exception)
		{
			$uniqueViolationVarName = str_contains($exception->getMessage(), "uniq_identifier_name") ? "name" : "email";
			return $this->json([
				"error" => "someone with the same $uniqueViolationVarName already exists",
			]);
		}
	}
}
