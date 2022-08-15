/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        "basin-green": "var(--basin-green)"
      }
    },
  },
  plugins: [
    function ({ addVariant }) {
      addVariant('child', '& > *');
    }
  ],
}
