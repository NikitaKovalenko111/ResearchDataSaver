import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig(() => {
  return {
    plugins: [
      react({
        babel: {
          plugins: [['babel-plugin-react-compiler']],
        },
      }),
    ],
    server: {
      port: 3000,
      host: true,
      watch: {
        usePolling: true
      },
      esbuild: {
        target: "esnext",
        platform: "linux"
      },
    },
  }
})
