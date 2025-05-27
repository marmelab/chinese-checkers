<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250526124247 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Create a view to get games of an account easily from React Admin.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<SQL
CREATE VIEW accounts_games AS (
	SELECT online_player.account_id, games.*, games.name, games.status, games.winner_name
	FROM online_player
	INNER JOIN games ON online_player.game_uuid = games.uuid
);
SQL);
	}

	public function down(Schema $schema): void
	{
		$this->addSql("DROP VIEW accounts_games;");
	}
}
