<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250506074023 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add online player name.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            ALTER TABLE online_player ADD name VARCHAR(255) NOT NULL DEFAULT 'Player'
        SQL
		);
		$this->addSql("ALTER TABLE online_player ALTER COLUMN name DROP DEFAULT");
	}

	public function down(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            ALTER TABLE online_player DROP name
        SQL
		);
	}
}
