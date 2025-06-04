<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250603144912 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Store last move in online games.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql("ALTER TABLE games ADD COLUMN last_move JSONB NULL;");
	}

	public function down(Schema $schema): void
	{
		$this->addSql("ALTER TABLE games DROP COLUMN last_move;");
	}
}
