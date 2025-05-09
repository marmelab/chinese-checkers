import React from "react";
import {Layout, usePageTitle} from "../Layout";
import {Link} from "react-router-dom";

/**
 * Show an error page.
 */
export function ErrorView() {
	usePageTitle("Error");

	return (
		<Layout>
			<main className={"error"}>
				<p className={"center"}>Sorry, something went seriously wrong!</p>

				<Link role={"button"} to={"/app"}>Go back to home</Link>
			</main>
		</Layout>
	);
}
