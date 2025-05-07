<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20250507155214 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add games creation and update date.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            ALTER TABLE games ADD created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
        SQL
		);
		$this->addSql(<<<'SQL'
            ALTER TABLE games ADD updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
        SQL
		);
	}

	public function down(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            ALTER TABLE games DROP created_at
        SQL
		);
		$this->addSql(<<<'SQL'
            ALTER TABLE games DROP updated_at
        SQL
		);
	}
}
