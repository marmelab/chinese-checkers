<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class HomeController extends AbstractController
{
	#[Route("/", name: "home")]
	public function index(): Response
	{
		// Return the response, with the rendered game.
		return $this->render("home.html.twig");
	}
}
