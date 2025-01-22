/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        'midnight-light': 'var(--color-midnight-light)',
        'midnight': 'var(--color-midnight)',
        'midnight-dark': 'var(--color-midnight-dark)',

        'green-light': 'var(--color-green-light)',
        'green': 'var(--color-green)',
        'green-dark': 'var(--color-green-dark)',

        'snow-light': 'var(--color-snow-light)',
        'snow': 'var(--color-snow)',
        'snow-dark': 'var(--color-snow-dark)',

        'primary-light': 'var(--color-primary-light)',
      },
      fontFamily: {
        sans: ['Lato', 'sans-serif'],
      },
    },
  },
  plugins: [],
}

