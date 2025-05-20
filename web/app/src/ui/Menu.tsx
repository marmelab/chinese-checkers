import React from "react";
import { NavLink } from "react-router-dom";
import { GameController, House, SignIn, User } from "@phosphor-icons/react";
import "./Menu.css";
import { useAuthenticatedAccount } from "../storage/authentication";

export function Menu() {
	const isAuthenticated = !!useAuthenticatedAccount();

	return (
		<nav className="menu">
			<ul>
				<li>
					<NavLink to="/app" end>
						<House /> Home
					</NavLink>
				</li>
				<li>
					<NavLink to="/app/play">
						<GameController /> Play
					</NavLink>
				</li>
				<li>
					<NavLink to={"/app/account"}>
						{isAuthenticated ? (
							<>
								<User /> Account
							</>
						) : (
							<>
								<SignIn /> Log in
							</>
						)}
					</NavLink>
				</li>
			</ul>
		</nav>
	);
}
