<?php

namespace App\Dto;

use Symfony\Component\Validator\Constraints as Assert;

class NewAccount
{
	#[Assert\Length(max: 180)]
	#[Assert\NotBlank]
	public string $name;

	#[Assert\Length(max: 180)]
	#[Assert\NotBlank]
	#[Assert\Email]
	public string $email;

	#[Assert\NotBlank]
	public string $password;
}
