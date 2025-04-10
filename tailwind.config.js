/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'media', // Changed from 'class' to 'media' for system preference based dark mode
  content: [
    "./frontend/templates/**/*.{html,js}",
    "./frontend/assets/css/**/*.{html,js}",
    "./node_modules/flowbite/**/*.js"
  ],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['Inter', 'sans-serif'],
        'mono': ['Consolas', 'Monaco', 'Courier New', 'monospace'],
      },
      colors: {
        primary: {
          light: "#6366F1",
          DEFAULT: "#4F46E5",
          dark: "#4338CA",
        },
        light: {
          base: "#ffffff",
          secondary: "#f6f8fa",
          card: "#eceef0",
        },
        dark: {
          base: "#1f1f1f",
          secondary: "#181818",
          card: "#1f2937",
        },
      },
    },
  },
  plugins: [
    require('flowbite/plugin')
  ],
};