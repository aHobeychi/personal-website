/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'media', // Changed from 'class' to 'media' for system preference based dark mode
  content: [
    "./frontend/templates/**/*.{html,js}",
    "./frontend/assets/css/**/*.{html,js}"
  ],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['Inter', 'sans-serif'],
        'heading': ['Inter Tight', 'sans-serif'],
        'body': ['Inter', 'sans-serif'],
        'mono': ['Consolas', 'Monaco', 'Courier New', 'monospace'],
      },
      fontSize: {
        'sm': '0.750rem',    // --font-size-sm
        'base': '1rem',      // --font-size-base
        'xl': '1.333rem',    // --font-size-xl
        '2xl': '1.777rem',   // --font-size-2xl
        '3xl': '2.369rem',   // --font-size-3xl
        '4xl': '3.158rem',   // --font-size-4xl
        '5xl': '4.210rem',   // --font-size-5xl
      },
      fontWeight: {
        'normal': 400,       // --font-weight-normal
        'bold': 700,         // --font-weight-bold
      },
      maxWidth: {
        'prose': '80ch',
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
    // Removed Flowbite plugin
  ],
};