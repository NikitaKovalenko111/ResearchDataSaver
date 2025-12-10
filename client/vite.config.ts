import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')

  return {
    define: {
      API_SERVER: JSON.stringify(env.API_SERVER)
    },
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
