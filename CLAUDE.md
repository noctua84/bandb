# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go web application for a Bed & Breakfast reservation system ("Fort Tranquility B&B"). Based on a Udemy course but with a custom modular template approach. Two rooms: General's Quarters and Major's Suite.

## Build & Run Commands

```bash
# Build and run the server (listens on :8080)
go build -o bandb ./cmd/web/*.go && ./bandb serve

# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test ./cmd/web/
go test ./src/handlers/
go test ./src/render/

# Run a single test
go test -run TestReservationFlow ./src/handlers/

# Build only (no run)
go build -v ./...

# Test coverage
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```

## Architecture

### Dependency Flow

`main.go` creates `AppConfig` → passes to `render`, `handlers`, `helpers` via package-level `New*()` init functions. All packages share the same `AppConfig` instance.

### Key Packages

| Package | Path | Purpose |
|---------|------|---------|
| **main** | `cmd/web/` | Entry point, routes, middleware, server startup |
| **config** | `src/config/` | `AppConfig` struct — central config passed to all packages |
| **handlers** | `src/handlers/` | HTTP handlers using Repository pattern (`Repo *Repository`) |
| **render** | `src/render/` | Template rendering with optional caching (`UseTemplate()`) |
| **forms** | `src/forms/` | Form validation (Required, MinLength, IsEmail) with error map |
| **helpers** | `src/helpers/` | `ClientError()` and `ServerError()` HTTP error helpers |
| **models** | `models/` | `Reservation` data struct and `TemplateData` for template rendering |

### Repository Pattern (Handlers)

All handlers are methods on `*Repository` which wraps `*AppConfig`. A global `Repo` variable is set during init via `NewHandlers()`. This enables dependency injection for testing.

```
handlers.NewRepo(appConfig) → repo
handlers.NewHandlers(repo)  → sets global Repo
```

### Template System

Templates use a three-tier naming convention in `templates/`:
- **`*.page.tmpl`** — Page content (defines `"content"` block)
- **`*.layout.tmpl`** — Base HTML layout (`base.layout.tmpl`)
- **`partials/*.tmpl`** — Reusable components (nav, bootstrap, sweetalert, custom-css)

`render.CreateTemplateCache()` builds a `map[string]*template.Template` keyed by page name without `.tmpl` extension (e.g., `"home.page"`). When `UseCache=false` (development), templates are rebuilt on every request.

`render.AddDefaultData()` injects CSRF token and session flash/warning/error messages into every template render.

### Middleware Stack (applied in order via chi)

1. `chi/middleware.Recoverer` — Panic recovery
2. `chi/middleware.Logger` — Request logging
3. `NoSurf` — CSRF protection (justinas/nosurf)
4. `SessionLoad` — Session load/save per request (scs/v2, 24h lifetime)

### Session

Uses `scs/v2` with cookie-based storage. `models.Reservation` is registered with `gob` for session serialization. Flash messages use `Pop` (one-time read).

## Testing Patterns

- **Unit tests** (`handlers_test.go`, `render_test.go`): Use `httptest.NewRequest` + `httptest.NewRecorder`, load session context manually
- **Integration tests** (`*_integration_test.go`): Use `httptest.NewTLSServer` with full middleware stack
- **Session flow tests** (`reservation_integration_test.go`): Use `cookiejar` for session persistence across requests, disable redirect following with `CheckRedirect`
- **Table-driven tests**: `handlerTest` struct with name, URL, method, expected status code
- **Test setup**: Each test package has `setup_test.go` with `TestMain` for shared initialization; handler tests use `getTestRepository()` and `getTestRoutes()`
- **Template path**: Handler integration tests reference templates at `../../templates` (relative to `src/handlers/`)

## Dependencies

- **chi/v5** — HTTP router
- **scs/v2** — Session management
- **nosurf** — CSRF protection
- Go 1.24.2

## CI

GitHub Actions (`.github/workflows/go.yml`): builds and tests on push/PR to main using `ubuntu-latest` and Go 1.24.2.
