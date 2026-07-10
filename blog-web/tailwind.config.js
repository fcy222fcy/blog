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
        primary: {
          DEFAULT: '#1B365D',
          darker: '#202A44',
          rgb: '27, 54, 93',
        },
        accent: '#1B365D',
        success: {
          DEFAULT: '#10b981',
          darker: '#059669',
          rgb: '16, 185, 129',
        },
        danger: {
          DEFAULT: '#e74c3c',
          rgb: '231, 76, 60',
        },
        warning: {
          DEFAULT: '#f59e0b',
          rgb: '245, 158, 11',
        },
      }
    },
  },
  plugins: [],
}
