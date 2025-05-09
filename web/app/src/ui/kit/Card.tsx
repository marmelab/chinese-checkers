import React from "react";
import {clsx} from "clsx";

/**
 * A generic card component.
 */
export function Card({
											 className,
											 ...props
										 }: React.DetailedHTMLProps<React.HTMLAttributes<HTMLDivElement>, HTMLDivElement>) {
	return (
		<div className={clsx("card", className)} {...props} />
	)
}
