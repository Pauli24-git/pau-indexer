module.exports = {
  content: [
    "./src/**/*.{html,js}",
    "./node_modules/tw-elements/dist/js/**/*.js",
    "./build/*.html"
  ],
  plugins: [require("tw-elements/dist/plugin.cjs")],
  darkMode: "class"
};

