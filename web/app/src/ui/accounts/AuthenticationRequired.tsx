import React from "react";
import { Link } from "react-router-dom";
import { getAuthenticationState } from "../../storage/authentication";
import { UserPlus } from "@phosphor-icons/react";
import { LoginForm } from "./LoginForm";

export function AuthenticationRequired({
	message,
	children,
}: React.PropsWithChildren<{
	message?: string;
}>) {
	const authState = getAuthenticationState();

	if (!authState)
		return (
			<>
				<header>
					<h1>Authentication required</h1>
				</header>

				<main className="error">
					{message && <p className="center">{message}</p>}

					<LoginForm />

					<hr />

					<p>No account?</p>

					<Link role="button" to={"/app/sign-up"}>
						<UserPlus /> Sign up
					</Link>
				</main>
			</>
		);

	return children;
}
