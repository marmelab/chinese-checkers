<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

final class Version20250505132418 extends AbstractMigration
{
	public function getDescription(): string
	{
		return "Add an online player, associated with the game.";
	}

	public function up(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            CREATE TABLE online_player (uuid VARCHAR(255) NOT NULL DEFAULT gen_random_uuid(), game_player INT NOT NULL, game_uuid UUID DEFAULT NULL, PRIMARY KEY(uuid))
        SQL
		);
		$this->addSql(<<<'SQL'
            CREATE INDEX IDX_60446BCB277CF5A7 ON online_player (game_uuid)
        SQL
		);
		$this->addSql(<<<'SQL'
            COMMENT ON COLUMN online_player.game_uuid IS '(DC2Type:uuid)'
        SQL
		);
		$this->addSql(<<<'SQL'
            ALTER TABLE online_player ADD CONSTRAINT FK_60446BCB277CF5A7 FOREIGN KEY (game_uuid) REFERENCES games (uuid) NOT DEFERRABLE INITIALLY IMMEDIATE
        SQL
		);
	}

	public function down(Schema $schema): void
	{
		$this->addSql(<<<'SQL'
            ALTER TABLE online_player DROP CONSTRAINT FK_60446BCB277CF5A7
        SQL
		);
		$this->addSql(<<<'SQL'
            DROP TABLE online_player
        SQL
		);
	}
}
