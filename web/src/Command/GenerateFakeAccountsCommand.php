<?php

namespace App\Command;

use App\Entity\Account;
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
	name: "app:generate-fake-accounts",
	description: "Generate fake accounts.",
)]
class GenerateFakeAccountsCommand extends Command
{
	public function __construct(
		private readonly EntityManagerInterface $entityManager,
		private readonly UserPasswordHasherInterface $passwordHasher,
	)
	{
		parent::__construct();
	}

	protected function configure(): void
	{
		$this
			->addOption("count", null, InputOption::VALUE_OPTIONAL, "Count of accounts to generate.");
	}

	protected function execute(InputInterface $input, OutputInterface $output): int
	{
		$io = new SymfonyStyle($input, $output);

		$faker = Faker\Factory::create();

		$count = 20;
		if (!empty($input->getOption("count")))
			$count = intval($input->getOption("count"));

		$io->info("Generating $count accounts...");

		foreach (range(0, $count) as $i)
		{
			$fakeAccount = new Account();
			$fakeAccount->setName($faker->name());
			$io->note("Creating {$fakeAccount->getName()}...");
			$fakeAccount->setEmail($faker->email());
			$fakeAccount->setPassword($this->passwordHasher->hashPassword($fakeAccount, $faker->password()));
			$this->entityManager->persist($fakeAccount);
		}

		$this->entityManager->flush();

		$io->success("$count accounts generated.");

		return Command::SUCCESS;
	}
}
