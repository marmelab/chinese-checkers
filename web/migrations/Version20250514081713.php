<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250514081713 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add accounts.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            CREATE TABLE "accounts" (id SERIAL NOT NULL, name VARCHAR(180) NOT NULL, email VARCHAR(180) NOT NULL, roles JSON NOT NULL, password VARCHAR(255) NOT NULL, PRIMARY KEY(id))
        SQL
		);
		$this->addSql(<<<'SQL'
            CREATE UNIQUE INDEX UNIQ_IDENTIFIER_NAME ON "accounts" (name)
        SQL
		);
		$this->addSql(<<<'SQL'
            CREATE UNIQUE INDEX UNIQ_IDENTIFIER_EMAIL ON "accounts" (email)
        SQL
		);
	}

	public function down(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            DROP TABLE "accounts"
        SQL
		);
	}
}
