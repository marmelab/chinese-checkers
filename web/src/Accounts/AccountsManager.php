<?php

namespace App\Accounts;

use App\Dto\NewAccount;
use App\Entity\Account;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\PasswordHasher\Hasher\UserPasswordHasherInterface;

readonly class AccountsManager
{
	public function __construct(
		private EntityManagerInterface      $entityManager,
		private UserPasswordHasherInterface $passwordHasher,
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
}
