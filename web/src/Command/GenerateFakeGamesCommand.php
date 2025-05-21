<?php

namespace App\Command;

use App\Entity\Account;
use App\Entity\Game;
use App\Entity\GamePlayer;
use App\Entity\OnlinePlayer;
use App\Game\BoardUtilities;
use Faker;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;
use Symfony\Component\Console\Style\SymfonyStyle;
use Symfony\Component\PasswordHasher\Hasher\UserPasswordHasherInterface;

#[AsCommand(
	name: "app:generate-fake-games",
	description: "Generate fake games.",
)]
class GenerateFakeGamesCommand extends Command
{
	private readonly Faker\Generator $faker;

	public function __construct(
		private readonly EntityManagerInterface $entityManager,
		private readonly BoardUtilities $boardUtilities,
	)
	{
		parent::__construct();
		$this->faker = Faker\Factory::create();
	}

	protected function configure(): void
	{
		$this
			->addOption("ongoing-games", null, InputOption::VALUE_OPTIONAL, "Count of ongoing games to generate.")
			->addOption("pending-games", null, InputOption::VALUE_OPTIONAL, "Count of pending (not fully joined) games to generate.")
			->addOption("won-games", null, InputOption::VALUE_OPTIONAL, "Count of won games to generate.");
	}

	protected function getShuffledBoard(): array
	{
		$board = $this->boardUtilities->getDefaultGameBoard();
		foreach ($board as &$row)
			shuffle($row);
		return $board;
	}

	protected function getGreenWinBoard(): array
	{
		// Reverse the default board: the 2 players have a winning state (should be impossible in real life).
		$board = array_reverse(
			array_map(fn (array $row) => array_reverse($row), $this->boardUtilities->getDefaultGameBoard())
		);
		// Change the first row to prevent the red from winning.
		$board[0][0] = 0;
		$board[0][5] = 2;
		return $board;
	}

	protected function addOnlinePlayer(Game $fakeGame, Account $account, GamePlayer $gamePlayer): void
	{
		$fakePlayer = new OnlinePlayer();
		$fakePlayer->setName($account->getName());
		$fakePlayer->setGamePlayer($gamePlayer);
		$fakePlayer->setAccount($account);
		$fakePlayer->setGame($fakeGame);
		$fakeGame->getPlayers()->add(
			$fakePlayer
		);
	}

	protected function getJoinCode(): string
	{
		return strtoupper(substr($this->faker->hexColor(), 1));
	}

	protected function execute(InputInterface $input, OutputInterface $output): int
	{
		$io = new SymfonyStyle($input, $output);

		$io->info("Retrieving all registered accounts to generate games for them.");
		$accounts = $this->entityManager->getRepository(Account::class)->findAll();

		{
			$count = 20;
			if (!empty($input->getOption("ongoing-games")))
				$count = intval($input->getOption("ongoing-games"));

			$io->info("Generating $count ongoing games...");

			foreach (range(1, $count) as $i)
			{
				$fakeGame = new Game();
				$fakeGame->setBoard($this->getShuffledBoard());
				$fakeGame->setCurrentPlayer(GamePlayer::random());
				$fakeGame->setJoinCode($this->getJoinCode());

				$this->addOnlinePlayer($fakeGame, $accounts[array_rand($accounts)], GamePlayer::Green);
				$this->addOnlinePlayer($fakeGame, $accounts[array_rand($accounts)], GamePlayer::Red);

				$this->entityManager->persist($fakeGame);
			}

			$this->entityManager->flush();

			$io->success("$count ongoing games generated.");
		}

		{
			$count = 5;
			if (!empty($input->getOption("pending-games")))
				$count = intval($input->getOption("pending-games"));

			$io->info("Generating $count pending games...");

			foreach (range(1, $count) as $i)
			{
				$fakeGame = new Game();
				$fakeGame->setBoard($this->getShuffledBoard());
				$fakeGame->setCurrentPlayer(GamePlayer::random());
				$fakeGame->setJoinCode($this->getJoinCode());

				$this->addOnlinePlayer($fakeGame, $accounts[array_rand($accounts)], GamePlayer::Green);

				$this->entityManager->persist($fakeGame);
			}

			$this->entityManager->flush();

			$io->success("$count pending games generated.");
		}

		{
			$count = 5;
			if (!empty($input->getOption("won-games")))
				$count = intval($input->getOption("won-games"));

			$io->info("Generating $count won games...");

			foreach (range(1, $count) as $i)
			{
				$fakeGame = new Game();
				$fakeGame->setBoard($this->getGreenWinBoard());
				$fakeGame->setCurrentPlayer(GamePlayer::Red);
				$fakeGame->setJoinCode($this->getJoinCode());
				$fakeGame->setWinner(GamePlayer::Green);

				$this->addOnlinePlayer($fakeGame, $accounts[array_rand($accounts)], GamePlayer::Green);
				$this->addOnlinePlayer($fakeGame, $accounts[array_rand($accounts)], GamePlayer::Red);

				$this->entityManager->persist($fakeGame);
			}

			$this->entityManager->flush();

			$io->success("$count won games generated.");
		}

		return Command::SUCCESS;
	}
}
