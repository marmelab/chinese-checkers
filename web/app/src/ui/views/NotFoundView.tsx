import React from "react";
import {Link} from "react-router-dom";
import {usePageTitle} from "../Layout";

export function NotFoundView() {
	usePageTitle("Not found");

	return (
		<main className={"not-found"}>
			<p className={"center"}>
				Sorry, we couldn't find what you were looking for!
			</p>

			<Link role={"button"} to={"/app"}>
				Go back to home
			</Link>
		</main>
	);
}
