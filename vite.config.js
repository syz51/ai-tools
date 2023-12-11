import { defineConfig } from "vite";

export default defineConfig(({ mode }) => {
  return {
    build: {
      minify: mode === "production",
      rollupOptions: {
        input: {
          main: "styles/main.css",
        },
        output: {
          assetFileNames: "main.css",
        },
      },
      outDir: "static",
      assetsDir: "",
      emptyOutDir: false,
    },
  };
});
