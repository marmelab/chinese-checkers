import React from "react";
import {Link} from "react-router-dom";
import {Layout} from "../Layout";

export function ErrorView() {
	return (
		<Layout>
			<header>
				<h1>Error</h1>
			</header>

			<main className={"error"}>
				<p className={"center"}>Sorry, something went seriously wrong!</p>

				<Link role={"button"} to={"/app"}>
					Go back to home
				</Link>
			</main>
		</Layout>
	);
}
