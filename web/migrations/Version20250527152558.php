<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250527152558 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Fix accounts games view when the same player is green and red.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<SQL
CREATE OR REPLACE VIEW accounts_games AS (
	SELECT DISTINCT online_player.account_id, games.*, games.name, games.status, games.winner_name
	FROM online_player
	INNER JOIN games ON online_player.game_uuid = games.uuid
);
SQL);
	}

	public function down(Schema $schema): void
	{
		$this->addSql(<<<SQL
CREATE OR REPLACE VIEW accounts_games AS (
	SELECT online_player.account_id, games.*, games.name, games.status, games.winner_name
	FROM online_player
	INNER JOIN games ON online_player.game_uuid = games.uuid
);
SQL);
	}
}
