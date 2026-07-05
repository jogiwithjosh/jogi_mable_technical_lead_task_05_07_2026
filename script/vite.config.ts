import { defineConfig } from "vite";
import dts from "vite-plugin-dts";

export default defineConfig({
    build: {
        lib: {
            entry: "src/index.ts",
            name: "TrackingSDK",
            fileName: "index",
            formats: ["es", "cjs"]
        },
        sourcemap: true
    },

    plugins: [
        dts()
    ]
});