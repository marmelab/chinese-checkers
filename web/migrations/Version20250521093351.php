<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250521093351 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add a computed field for game name.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
CREATE FUNCTION name(games) RETURNS TEXT LANGUAGE SQL AS $$
  SELECT COALESCE((ARRAY_AGG(accounts.name))[1], 'Green') || ' VS ' || COALESCE((ARRAY_AGG(accounts.name))[2], 'Red')
	FROM online_player
		INNER JOIN accounts ON online_player.account_id = accounts.id
	WHERE online_player.game_uuid = $1.uuid
	GROUP BY online_player.game_uuid
$$;
SQL);
	}

	public function down(Schema $schema): void
	{
		$this->addSql("DROP FUNCTION name(games);");
	}
}
