<?php

namespace App\Accounts;

use App\Dto\NewUser;
use App\Entity\User;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\PasswordHasher\Hasher\UserPasswordHasherInterface;

readonly class UsersManager
{
	public function __construct(
		private EntityManagerInterface      $entityManager,
		private UserPasswordHasherInterface $passwordHasher,
	) {}

	public function create(NewUser $newUser): User
	{
		$user = new User();

		$user->setName(trim($newUser->name));
		$user->setEmail($newUser->email);
		$user->setPassword($this->passwordHasher->hashPassword($user, $newUser->password));

		$this->entityManager->persist($user);
		$this->entityManager->flush();

		return $user;
	}
}
