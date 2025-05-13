<?php

namespace App\Repository;

use App\Entity\Game;
use Doctrine\Bundle\DoctrineBundle\Repository\ServiceEntityRepository;
use Doctrine\Persistence\ManagerRegistry;

/**
 * Games repository.
 */
class GamesRepository extends ServiceEntityRepository
{
	public function __construct(ManagerRegistry $registry)
	{
		parent::__construct($registry, Game::class);
	}

	/**
	 * Find all ongoing (2 players, not finished) games.
	 * @return Game[]
	 */
	public function findAllOngoingGames(): array
	{
		return $this->getEntityManager()->createQueryBuilder()

			->select("games")
			->from($this->getEntityName(), "games")
			->where("(SELECT COUNT(players) FROM App\Entity\OnlinePlayer players WHERE players.game = games.uuid) = 2")
			->andWhere("games.winner IS NULL")
			->orderBy("games.updatedAt", "DESC")
			->addOrderBy("games.createdAt", "DESC")

			->getQuery()->execute();
	}
}
