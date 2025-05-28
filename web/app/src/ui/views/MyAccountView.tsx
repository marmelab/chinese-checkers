import React from "react";
import { SignOut } from "@phosphor-icons/react";
import { AuthenticationRequired } from "../accounts/AuthenticationRequired";
import {
	clearAuthenticatedAccount,
	useAuthenticatedAccount,
} from "../../storage/authentication";
import { useNavigate } from "react-router-dom";
import { useFetchAuthenticatedAccount } from "../../api/accounts";

export function MyAccountView() {
	return (
		<AuthenticationRequired>
			<MyAccount />
		</AuthenticationRequired>
	);
}

export function MyAccount() {
	const navigate = useNavigate();

	const authenticatedAccountQuery = useFetchAuthenticatedAccount();

	return (
		<>
			<header>
				<h1>My Account</h1>
			</header>
			<main className="account">
				<p className="center">
					Connected as <strong>{authenticatedAccountQuery.data.name}</strong>
				</p>

				<button
					className="log-out"
					onClick={() => {
						clearAuthenticatedAccount();
						navigate("/app");
					}}
				>
					<SignOut /> Log out
				</button>
			</main>
		</>
	);
}
