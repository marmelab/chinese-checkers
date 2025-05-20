<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250520090838 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add the related account of an online player.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            ALTER TABLE online_player ADD account_id INT DEFAULT NULL
        SQL
		);
		$this->addSql(<<<'SQL'
            ALTER TABLE online_player ADD CONSTRAINT FK_60446BCB9B6B5FBA FOREIGN KEY (account_id) REFERENCES "accounts" (id) NOT DEFERRABLE INITIALLY IMMEDIATE
        SQL
		);
		$this->addSql(<<<'SQL'
            CREATE INDEX IDX_60446BCB9B6B5FBA ON online_player (account_id)
        SQL
		);
	}

	public function down(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            ALTER TABLE online_player DROP CONSTRAINT FK_60446BCB9B6B5FBA
        SQL
		);
		$this->addSql(<<<'SQL'
            DROP INDEX IDX_60446BCB9B6B5FBA
        SQL
		);
		$this->addSql(<<<'SQL'
            ALTER TABLE online_player DROP account_id
        SQL
		);
	}
}
