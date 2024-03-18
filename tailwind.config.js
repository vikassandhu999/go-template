const colors = require("tailwindcss/colors");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["internal/controller/**/*.templ"],
  theme: {
    extend: {
      colors: {
        primary: colors.blue,
        secondary: colors.yellow,
        neutral: colors.gray,
      },
    },
  },
  plugins: [require("flowbite/plugin")],
};
