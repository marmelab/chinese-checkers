<?php

namespace App\Repository;

use App\Entity\Account;
use Doctrine\Bundle\DoctrineBundle\Repository\ServiceEntityRepository;
use Doctrine\Persistence\ManagerRegistry;
use Symfony\Bridge\Doctrine\Security\User\UserLoaderInterface;
use Symfony\Component\Security\Core\Exception\UnsupportedUserException;
use Symfony\Component\Security\Core\User\PasswordAuthenticatedUserInterface;
use Symfony\Component\Security\Core\User\PasswordUpgraderInterface;
use Symfony\Component\Security\Core\User\UserInterface;

/**
 * @extends ServiceEntityRepository<Account>
 */
class AccountsRepository extends ServiceEntityRepository implements PasswordUpgraderInterface, UserLoaderInterface
{
	public function __construct(ManagerRegistry $registry)
	{
		parent::__construct($registry, Account::class);
	}

	/**
	 * Used to upgrade (rehash) the account's password automatically over time.
	 */
	public function upgradePassword(PasswordAuthenticatedUserInterface $account, string $newHashedPassword): void
	{
		if (!$account instanceof Account)
		{
			throw new UnsupportedUserException(sprintf('Instances of "%s" are not supported.', $account::class));
		}

		$account->setPassword($newHashedPassword);
		$this->getEntityManager()->persist($account);
		$this->getEntityManager()->flush();
	}

	public function loadUserByIdentifier(string $usernameOrEmail): ?UserInterface
	{
		return $this->getEntityManager()->createQuery(
			"SELECT accounts
			FROM App\Entity\Account accounts
			WHERE accounts.name = :query
			OR accounts.email = :query
			"
		)
			->setParameter("query", $usernameOrEmail)
			->getOneOrNullResult();
	}
}
