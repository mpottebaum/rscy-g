import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  build: {
    manifest: true,
    rollupOptions: {
      input: './src/gates-of-rsc.tsx',
      output: {
        entryFileNames: 'js/gates-of-rsc.js'
      }
    },
  }
})
