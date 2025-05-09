import React from "react";

/**
 * Loader component.
 */
export function Loader() {
	return (
		<div className={"loader"}></div>
	);
}

/**
 * Loader view component.
 */
export function LoaderView() {
	return (
		<main className={"loader"}>
			<Loader/>
		</main>
	);
}
