# Deployment readiness TODO

- [ ] Backend: add config/env loading (PORT, JWT_SECRET validation, DB_SSLMODE, CORS_ORIGINS, MIGRATE_ON_STARTUP)
- [ ] Backend: replace router.Run with http.Server + graceful shutdown + timeouts
- [ ] Backend: add health endpoint(s) (/healthz, optional /readyz)
- [ ] Backend: add CORS middleware configured from env
- [ ] Backend: secure websocket origin check using allowed origins from env
- [ ] Backend: make AutoMigrate optional via MIGRATE_ON_STARTUP
- [ ] Backend: (optional) ensure production logging/error responses are consistent
- [ ] Frontend: make API baseURL configurable via VITE_API_BASE_URL (default /api)
- [ ] Backend: add Dockerfile (multi-stage, run non-root)
- [ ] Frontend: add Dockerfile (build + serve static)
- [ ] Add docker-compose.yml (frontend + backend + postgres)
- [ ] Add .env.example (document required env vars)
- [x] Update README.md with deployment instructions (single-domain reverse proxy + separate-domain CORS notes)
- [x] Add nginx reverse proxy config (optional) so frontend + backend work on one domain


- [ ] Smoke test: docker compose up --build; verify /healthz and frontend login flow

