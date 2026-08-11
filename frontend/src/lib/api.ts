import axios from 'axios'

// In dev, '/api' is proxied to localhost:8080 by vite.config.ts. GitHub
// Pages is a static host with no proxy and no backend of its own, so in
// production this must point at wherever your Go backend is actually
// deployed (Render, Fly.io, a VPS, etc.) — set VITE_API_URL at build time.
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
