import React, { useState } from "react";
import { SignIn } from "@phosphor-icons/react";
import { showErrorToast } from "../showErrorToast";
import { useNavigate } from "react-router-dom";
import { authenticate, InvalidCredentialsError } from "../../api/accounts";

export function LoginForm({ redirectTo }: { redirectTo?: string }) {
	const navigate = useNavigate();

	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");

	const [error, setError] = useState("");

	return (
		<form
			onSubmit={async (event) => {
				event.preventDefault();
				try {
					await authenticate(username, password);
					navigate((redirectTo ?? 0) as any);
				} catch (error) {
					if (error instanceof InvalidCredentialsError) {
						setError("Invalid credentials.");
					} else {
						showErrorToast(error);
					}
					setPassword("");
				}
			}}
		>
			<label>
				Username (or email)
				<input
					type="text"
					name="username"
					value={username}
					onChange={(event) => setUsername(event.currentTarget.value)}
				/>
				{error && <span className="error">{error}</span>}
			</label>

			<label>
				Password
				<input
					type="password"
					name="password"
					value={password}
					onChange={(event) => setPassword(event.currentTarget.value)}
				/>
			</label>

			<button>
				<SignIn /> Log in
			</button>
		</form>
	);
}
