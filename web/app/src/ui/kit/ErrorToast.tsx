import React from "react";
import "./ErrorToast.css";
import { Toast } from "./Toast";

export function ErrorToast({ children }: { children: string }) {
	return (
		<Toast className="error">
			<p>{children}</p>
		</Toast>
	);
}
