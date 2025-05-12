import React from "react";
import { createRoot } from "react-dom/client";
import { App } from "./ui/App";

document.addEventListener("DOMContentLoaded", () => {
	const root = createRoot(document.body);
	root.render(<App />);
});
