<?php

namespace App\Accounts;

use App\Dto\NewAccount;
use App\Entity\Account;
use Doctrine\ORM\EntityManagerInterface;
use Lexik\Bundle\JWTAuthenticationBundle\Services\JWTTokenManagerInterface;
use Symfony\Component\HttpFoundation\Cookie;
use Symfony\Component\PasswordHasher\Hasher\UserPasswordHasherInterface;
use Symfony\Component\Security\Core\User\UserInterface;

readonly class AccountsManager
{
	public function __construct(
		private EntityManagerInterface      $entityManager,
		private UserPasswordHasherInterface $passwordHasher,
		private JWTTokenManagerInterface    $jwtManager,
	) {}

	public function create(NewAccount $newAccount): Account
	{
		$account = new Account();

		$account->setName(trim($newAccount->name));
		$account->setEmail($newAccount->email);
		$account->setPassword($this->passwordHasher->hashPassword($account, $newAccount->password));

		$this->entityManager->persist($account);
		$this->entityManager->flush();

		return $account;
	}

	/**
	 * @param UserInterface $account
	 * @return bool
	 */
	public function isAdmin(UserInterface $account): bool
	{
		return in_array("ROLE_ADMIN", $account->getRoles());
	}

	public function getAuthenticationToken(UserInterface $account): string
	{
		return $this->jwtManager->createFromPayload($account, [
			"role" => $this->isAdmin($account) ? "admin" : null,
		]);
	}

	public function getAuthenticationCookie(UserInterface $account): Cookie
	{
		return (
			Cookie::create("authentication")
				->withValue($this->getAuthenticationToken($account))
				->withSecure()
				->withHttpOnly(false)
		);
	}
}
