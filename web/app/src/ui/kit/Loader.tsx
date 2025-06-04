import React from "react";
import { clsx } from "clsx";
import "./Loader.css";

export function Loader({ className }: { className?: string }) {
	return <div className={clsx("loader", className)}></div>;
}

export function LoaderView() {
	return (
		<main className="loader">
			<Loader />
		</main>
	);
}
