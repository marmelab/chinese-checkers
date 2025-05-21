<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250520152149 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add a computed field for current game status.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
CREATE FUNCTION status(games) RETURNS TEXT LANGUAGE SQL AS $$
    SELECT
		CASE WHEN ($1.winner IS NOT NULL)
			THEN 'finished'
		ELSE
			CASE WHEN (SELECT COUNT(*) FROM online_player WHERE game_uuid = $1.uuid) >= 2 THEN 'started' ELSE 'pending' END
		END
	FROM games
$$;
SQL);
	}

	public function down(Schema $schema): void
	{
		$this->addSql("DROP FUNCTION status(games);");
	}
}
