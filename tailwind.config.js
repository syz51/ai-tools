import daisyui from "daisyui";
import colors from "tailwindcss/colors";

/** @type {import('tailwindcss').Config} */
export default {
  content: ["**/*.{html,templ}"],
  plugins: [daisyui],
  theme: {
    extend: {
      colors: {
        neutralgray: colors.neutral,
      },
    },
  },
};
