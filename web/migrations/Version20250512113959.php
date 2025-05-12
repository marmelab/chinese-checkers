<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250512113959 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add game join code.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            ALTER TABLE games ADD join_code VARCHAR(255) DEFAULT NULL
        SQL
		);
		$this->addSql(<<<'SQL'
            CREATE UNIQUE INDEX UNIQ_FF232B31E64D7D01 ON games (join_code)
        SQL
		);
	}

	public function down(Schema $schema): void
	{

		$this->addSql(<<<'SQL'
            ALTER TABLE games DROP join_code
        SQL
		);
	}
}
