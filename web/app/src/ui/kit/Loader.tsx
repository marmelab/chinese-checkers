import React from "react";
import "./Loader.css";

export function Loader() {
	return <div className={"loader"}></div>;
}

export function LoaderView() {
	return (
		<main className={"loader"}>
			<Loader />
		</main>
	);
}
