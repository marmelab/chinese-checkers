import React from "react";
import {clsx} from "clsx";
import "./Card.css";

export function Card({
	className,
	...props
}: React.DetailedHTMLProps<
	React.HTMLAttributes<HTMLDivElement>,
	HTMLDivElement
>) {
	return <div className={clsx("card", className)} {...props} />;
}
