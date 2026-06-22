/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        primary: '#117361',
        'primary-dark': '#0d5a4c',
        accent: '#007bff',
      }
    },
  },
  plugins: [],
}
