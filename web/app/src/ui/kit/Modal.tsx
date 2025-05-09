import React from "react";
import {clsx} from "clsx";

export function Modal({
	className,
	...props
}: React.DetailedHTMLProps<
	React.HTMLAttributes<HTMLDivElement>,
	HTMLDivElement
>) {
	return <div className={clsx("modal", className)} {...props} />;
}
