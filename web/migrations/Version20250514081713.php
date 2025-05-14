<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250514081713 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add users.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            CREATE TABLE "users" (id SERIAL NOT NULL, name VARCHAR(180) NOT NULL, email VARCHAR(180) NOT NULL, roles JSON NOT NULL, password VARCHAR(255) NOT NULL, PRIMARY KEY(id))
        SQL
		);
		$this->addSql(<<<'SQL'
            CREATE UNIQUE INDEX UNIQ_IDENTIFIER_NAME ON "users" (name)
        SQL
		);
		$this->addSql(<<<'SQL'
            CREATE UNIQUE INDEX UNIQ_IDENTIFIER_EMAIL ON "users" (email)
        SQL
		);
	}

	public function down(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            DROP TABLE "users"
        SQL
		);
	}
}
