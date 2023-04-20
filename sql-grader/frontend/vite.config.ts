import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [react()],
	base: '/',
	server: {
		proxy: {
			'/api': {
				target: process.env.VITE_BACKEND_URL,
				changeOrigin: true,
			},
		},
	},
})
