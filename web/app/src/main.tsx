import React from "react";
import {createRoot} from "react-dom/client";
import {App} from "./ui/app";

import "../styles/chinese-checkers.less";

document.addEventListener("DOMContentLoaded", () => {
	const root = createRoot(document.body);
	root.render(<App/>);
});
