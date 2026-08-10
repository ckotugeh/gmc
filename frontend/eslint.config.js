import js from "@eslint/js";

export default [
  // 1. Tell ESLint to ignore build artifacts and dependency directories
  {
    ignores: ["dist/**", "build/**", "node_modules/**"]
  },
  
  // 2. Main rule configurations for your actual code files
  js.configs.recommended,
  {
    languageOptions: {
      ecmaVersion: "latest",
      sourceType: "module",
      globals: {
        // Core browser context globals
        window: "readonly",
        document: "readonly",
        console: "readonly",
        process: "readonly",
        setTimeout: "readonly",
        clearTimeout: "readonly",
        fetch: "readonly",
        FormData: "readonly",
        MutationObserver: "readonly",
        AbortController: "readonly",
      },
    },
    rules: {
      // Prevents unused variable warnings from throwing errors during dev cycles
      "no-unused-vars": "warn",
      "no-undef": "error",
    },
  },
];
