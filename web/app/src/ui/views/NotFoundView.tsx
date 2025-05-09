import React from "react";
import {usePageTitle} from "../Layout";
import {Link} from "react-router-dom";

export function NotFoundView() {
	usePageTitle("Not found");

	return (
		<main className={"not-found"}>
			<p className={"center"}>Sorry, we couldn't find what you were looking for!</p>

			<Link role={"button"} to={"/app"}>Go back to home</Link>
		</main>
	);
}
