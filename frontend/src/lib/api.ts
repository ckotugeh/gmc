import axios from 'axios'

// In dev: Vite proxies /api → localhost:8080 (vite.config.ts)
// In production: Render rewrites /api → backend service (render.yaml)
// VITE_API_URL defaults to /api so both work without code changes
const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || '/api',
})

// Attach JWT token to every request if present
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

export default api
