import { defineConfig } from "vite";

export default defineConfig({
  root: "./", // Specify the root directory
  publicDir: "public", // Specify the static file directory
  server: {
    port: 8080, // Specify the development server port
  },
});

