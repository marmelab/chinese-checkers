<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250505083131 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add game board table.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            CREATE TABLE games (uuid UUID NOT NULL DEFAULT gen_random_uuid(), board JSONB NOT NULL, current_player SMALLINT NOT NULL, PRIMARY KEY(uuid))
        SQL
		);
		$this->addSql(<<<'SQL'
            COMMENT ON COLUMN games.uuid IS '(DC2Type:uuid)'
        SQL
		);
	}

	public function down(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            DROP TABLE games
        SQL
		);
	}
}
