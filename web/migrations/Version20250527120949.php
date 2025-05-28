<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250527120949 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add roles for Postgrest API.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<SQL
CREATE ROLE authenticator LOGIN NOINHERIT NOCREATEDB NOCREATEROLE NOSUPERUSER;
SQL);
		$this->addSql(<<<SQL
CREATE ROLE admin NOLOGIN;
SQL);

		$this->addSql("GRANT admin TO authenticator;");

		$this->addSql("GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO admin;");
		$this->addSql("GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA public TO admin;");
		$this->addSql("ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO admin;");
		$this->addSql("ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT EXECUTE ON FUNCTIONS TO admin;");
	}

	public function down(Schema $schema): void
	{
		$this->addSql("ALTER DEFAULT PRIVILEGES IN SCHEMA public REVOKE ALL PRIVILEGES ON TABLES TO admin;");
		$this->addSql("ALTER DEFAULT PRIVILEGES IN SCHEMA public REVOKE EXECUTE ON FUNCTIONS TO admin;");

		$this->addSql("DROP ROLE authenticator;");
		$this->addSql("DROP ROLE admin;");
	}
}
