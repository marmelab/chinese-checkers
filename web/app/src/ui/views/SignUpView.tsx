import React from "react";
import { UserPlus } from "@phosphor-icons/react";

export function SignUpView() {
	return (
		<>
			<header>
				<h1>Sign up</h1>
			</header>

			<main className="sign-up">
				<p className="center">Create an account to join your friends games!</p>

				<form
					onSubmit={(event) => {
						event.preventDefault();
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
						/>
					</label>

					<label htmlFor="password">
						Password
						<input
							type="password"
							id="password"
							name="password"
							required={true}
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
