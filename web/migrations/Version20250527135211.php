<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250527135211 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add admin boolean to accounts.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
CREATE FUNCTION admin(accounts) RETURNS BOOLEAN LANGUAGE SQL AS $$
  SELECT COALESCE('ROLE_ADMIN' = ANY(ARRAY_AGG(JSON_ARRAY_ELEMENTS_TEXT)), FALSE)
  FROM JSON_ARRAY_ELEMENTS_TEXT($1.roles)
$$;
SQL);
	}

	public function down(Schema $schema): void
	{
		$this->addSql("DROP FUNCTION admin(accounts);");
	}
}
