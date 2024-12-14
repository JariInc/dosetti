/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/html/**/*.html"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["garden"],
  },
};
