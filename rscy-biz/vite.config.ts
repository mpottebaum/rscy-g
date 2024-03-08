import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react'
import { resolve } from 'path';
import replace from '@rollup/plugin-replace'

export default defineConfig({
  plugins: [
    react(),
    replace({
      'process.env': 'import.meta.env',
      preventAssignment: true,
    }),
  ],
  // define: {
  //     'process.env.NODE_ENV': '"production"',
  // },
  build: {
    lib: {
      entry: resolve(__dirname, 'src/gates-of-rsc.tsx'),
      name: 'GatesOfRsc',
      fileName: 'gates-of-rsc',
    },
  },
})
