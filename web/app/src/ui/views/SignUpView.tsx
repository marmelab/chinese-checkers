import React, { useState } from "react";
import { UserPlus } from "@phosphor-icons/react";
import { createAccount } from "../../api/accounts";
import { handleCallbackError } from "../CallbackErrorHandler";
import { useNavigate } from "react-router-dom";

export function SignUpView() {
	const navigate = useNavigate();

	const [name, setName] = useState("");
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");

	return (
		<>
			<header>
				<h1>Sign up</h1>
			</header>

			<main className="sign-up">
				<p className="center">Create an account to join your friends games!</p>

				<form
					onSubmit={async (event) => {
						event.preventDefault();
						try {
							await createAccount(name, email, password);
							navigate("/app/registered");
						} catch (error) {
							handleCallbackError(error);
						}
					}}
				>
					<label htmlFor="name">
						Name
						<input
							type="text"
							id="name"
							name="name"
							min={1}
							max={180}
							required={true}
							value={name}
							onChange={(event) => setName(event.currentTarget.value)}
						/>
					</label>

					<label htmlFor="email">
						Email
						<input
							type="email"
							id="email"
							name="email"
							min={1}
							max={180}
							required={true}
							value={email}
							onChange={(event) => setEmail(event.currentTarget.value)}
						/>
					</label>

					<label htmlFor="password">
						Password
						<input
							type="password"
							id="password"
							name="password"
							required={true}
							value={password}
							onChange={(event) => setPassword(event.currentTarget.value)}
						/>
					</label>

					<button>
						<UserPlus /> Sign up!
					</button>
				</form>
			</main>
		</>
	);
}
