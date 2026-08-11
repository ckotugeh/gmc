import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";

export default defineConfig({
  // GitHub Pages serves project sites at https://<user>.github.io/<repo>/,
  // so every asset URL needs that /<repo>/ prefix baked in at build time.
  // Replace REPO_NAME below with your actual GitHub repository name
  // (skip this / leave it as '/' only if this deploys to a user/org page
  // at the domain root, e.g. <user>.github.io).
  base: '/REPO_NAME/',
  plugins: [
    react(),
    tailwindcss(),
  ],
  server: {
    port: 5173,
    proxy: {
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  },
});
