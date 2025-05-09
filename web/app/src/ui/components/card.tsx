import React from "react";
import {classes} from "../../utils";

/**
 * A generic card component.
 */
export function Card({
											 className,
											 ...props
										 }: React.DetailedHTMLProps<React.HTMLAttributes<HTMLDivElement>, HTMLDivElement>) {
	return (
		<div className={classes("card", className)} {...props} />
	)
}
