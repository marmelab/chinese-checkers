<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250526144412 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Cascading updates and deletion of a game to online players.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<SQL
ALTER TABLE online_player
	DROP CONSTRAINT fk_60446bcb277cf5a7;
SQL);
		$this->addSql(<<<SQL
ALTER TABLE online_player
	ADD CONSTRAINT fk_60446bcb277cf5a7
		FOREIGN KEY (game_uuid) REFERENCES games
			ON UPDATE CASCADE ON DELETE CASCADE;
SQL);
	}

	public function down(Schema $schema): void
	{
		$this->addSql(<<<SQL
ALTER TABLE online_player
	DROP CONSTRAINT fk_60446bcb277cf5a7;
SQL);
		$this->addSql(<<<SQL
ALTER TABLE online_player
	ADD CONSTRAINT fk_60446bcb277cf5a7
		FOREIGN KEY (game_uuid) REFERENCES games;
SQL);
	}
}
