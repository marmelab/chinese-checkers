import React from "react";
import {Link} from "react-router-dom";
import {Layout, usePageTitle} from "../Layout";

export function ErrorView() {
	usePageTitle("Error");

	return (
		<Layout>
			<main className={"error"}>
				<p className={"center"}>Sorry, something went seriously wrong!</p>

				<Link role={"button"} to={"/app"}>
					Go back to home
				</Link>
			</main>
		</Layout>
	);
}
