import React from "react";
import { CheckCircle, SignIn } from "@phosphor-icons/react";
import "./RegisteredView.css";
import { Link } from "react-router-dom";

export function RegisteredView() {
	return (
		<>
			<header>
				<h1>Registered</h1>
			</header>
			<main className="registered">
				<CheckCircle weight="regular" size={"8em"} />
				<p className="center">
					You are now successfully registered to Chinese Checkers!
				</p>

				<Link role="button" to={"/app/account"}>
					<SignIn /> Log in
				</Link>
			</main>
		</>
	);
}
