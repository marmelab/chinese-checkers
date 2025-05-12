<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250512130819 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add game winner.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            ALTER TABLE games ADD COLUMN winner SMALLINT NULL DEFAULT NULL
        SQL
		);
	}

	public function down(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            ALTER TABLE games DROP COLUMN winner
        SQL
		);
	}
}
