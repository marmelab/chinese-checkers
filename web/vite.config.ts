import { ConfigEnv, defineConfig, UserConfig } from "vite";
import react from "@vitejs/plugin-react";
import path from "node:path";

// https://vitejs.dev/config/

export default defineConfig(({ mode }: ConfigEnv): UserConfig => {
	return {
		base: "/app",

		server: {
			proxy: {
				"^/(?!app)": "http://localhost",
			},
		},

		build: {
			sourcemap: true,
			outDir: path.resolve("public", "app"),
			minify: mode == "production" ? "esbuild" : false,
		},

		plugins: [react()],
	};
});
