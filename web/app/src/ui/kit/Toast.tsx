import React, { useEffect } from "react";
import { clsx } from "clsx";
import { useToastClose } from "./Toasts";

/**
 * In milliseconds.
 */
const AUTO_CLOSE_DELAY = 10_000;

export function Toast({
	className,
	onClick,
	...props
}: React.DetailedHTMLProps<
	React.HTMLAttributes<HTMLDivElement>,
	HTMLDivElement
>) {
	const close = useToastClose();

	useEffect(() => {
		setTimeout(close, AUTO_CLOSE_DELAY);
	}, []);

	return (
		<div
			className={clsx("toast", className)}
			onClick={(event) => {
				onClick?.(event);
				close();
			}}
			{...props}
		/>
	);
}
