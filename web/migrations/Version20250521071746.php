<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250521071746 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add a computed field for winner account name.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
CREATE FUNCTION winner_name(games) RETURNS TEXT LANGUAGE SQL AS $$
  SELECT accounts.name
	FROM online_player
		INNER JOIN accounts ON online_player.account_id = accounts.id
	WHERE online_player.game_uuid = $1.uuid AND online_player.game_player = $1.winner
$$;
SQL);
	}

	public function down(Schema $schema): void
	{
		$this->addSql("DROP FUNCTION winner_name(games);");
	}
}
