import React from "react";
import {Link} from "react-router-dom";
import "./NotFoundView.css";

export function NotFoundView() {
	return (
		<>
			<header>
				<h1>Not found</h1>
			</header>
			<main className={"not-found"}>
				<p className={"center"}>
					Sorry, we couldn't find what you were looking for!
				</p>

				<Link role={"button"} to={"/app"}>
					Go back to home
				</Link>
			</main>
		</>
	);
}
