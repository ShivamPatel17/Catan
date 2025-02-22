import { defineConfig } from "vite";
import path from "path";

export default defineConfig({
  root: "./", // Specify the root directory
  publicDir: "public", // Specify the static file directory
  server: {
    port: 5173, // Specify the development server port
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"), // Maps '@' to './src'
      utils: path.resolve(__dirname, "./src/utils"),
      config: path.resolve(__dirname, "./src/config"),
      assets: path.resolve(__dirname, "./src/assets"),
      game: path.resolve(__dirname, "./src/game"),
      builders: path.resolve(__dirname, "./src/game/messages/builders"),
    },
  },
});
