# Go Programming Course

A practical, developer-focused Go course built around explanation, coding exercises, code review, and progressively larger projects.

## Learning Method

Each lesson follows this structure:

1. Concept explanation
2. Small code examples
3. Exercises
4. Challenge prediction
5. Code review
6. Idiomatic Go improvements
7. Practical application

The course prioritizes writing code instead of only reading theory.

---

# Course Roadmap

## Phase 1 — Go Fundamentals

Goal: understand the language syntax, core data structures, functions, types, and error handling.

- [x] Go program structure
- [x] Packages and `main`
- [x] Variables and constants
- [x] Basic types
- [x] Functions and multiple return values
- [x] Conditionals
- [x] `switch`
- [x] Loops
- [x] Arrays
- [x] Slices
- [x] `range`
- [x] Maps
- [x] Comma-ok idiom
- [x] Structs
- [x] Methods
- [x] Value and pointer receivers
- [x] Interfaces
- [x] Type assertions
- [x] Type switches
- [x] Basic error handling
- [x] Error wrapping
- [x] Sentinel errors
- [x] `errors.Is`
- [x] `defer`
- [x] `panic` and `recover`
- [x] Package organization
- [x] Go modules
- [x] Exported and unexported identifiers
- [x] Working with strings and runes
- [x] Basic file handling
- [x] JSON encoding and decoding
- [ ] Phase 1 final project

### Phase 1 Final Project

Build a command-line task manager with:

- Task creation
- Task listing
- Task completion
- Task deletion
- Structs and methods
- Slices and maps
- Input validation
- Error handling
- JSON file persistence
- Multiple packages

---

## Phase 2 — Idiomatic Go and Project Structure

Goal: learn how professional Go projects are organized and how Go code is commonly written in real applications.

### Topics

- [ ] Go modules in depth
- [ ] `go.mod`
- [ ] `go.sum`
- [ ] Package design
- [ ] `internal` packages
- [ ] `cmd` directories
- [ ] Constructor functions
- [ ] Dependency injection
- [ ] Small interfaces
- [ ] Interface placement
- [ ] Error design
- [ ] Error wrapping
- [ ] Custom error types
- [ ] Configuration patterns
- [ ] Environment variables
- [ ] Formatting with `gofmt`
- [ ] Static analysis with `go vet`
- [ ] Basic linting
- [ ] Code organization and naming

### Phase 2 Project

Refactor the task manager into a structured application:

```text
go-task-manager/
├── cmd/
│   └── task/
│       └── main.go
├── internal/
│   ├── task/
│   ├── storage/
│   └── service/
├── go.mod
└── README.md
```

---

## Phase 3 — Testing

Goal: write reliable Go applications with automated tests.

### Topics

- [ ] The `testing` package
- [ ] Unit tests
- [ ] Table-driven tests
- [ ] Subtests
- [ ] Test helpers
- [ ] Testing errors
- [ ] Testing interfaces
- [ ] Mock and fake implementations
- [ ] HTTP handler testing
- [ ] Code coverage
- [ ] Benchmarks
- [ ] Race detection
- [ ] Fuzz testing

### Important Commands

```bash
go test ./...
go test -v ./...
go test -cover ./...
go test -race ./...
go test -bench=. ./...
go test -fuzz=Fuzz ./...
```

### Phase 3 Project

Add complete tests to the task manager:

- Service tests
- Repository tests
- Validation tests
- File storage tests
- Error-path tests
- Fuzz tests

---

## Phase 4 — HTTP APIs

Goal: build backend services using Go's standard library.

### Topics

- [ ] `net/http`
- [ ] HTTP handlers
- [ ] `http.Handler`
- [ ] `http.HandlerFunc`
- [ ] Routing
- [ ] Request methods
- [ ] Request headers
- [ ] Request bodies
- [ ] JSON requests
- [ ] JSON responses
- [ ] HTTP status codes
- [ ] Middleware
- [ ] Request validation
- [ ] Context
- [ ] Timeouts
- [ ] Graceful shutdown
- [ ] Structured logging
- [ ] API error responses

### Phase 4 Project

Convert the task manager into a REST API:

```text
POST   /tasks
GET    /tasks
GET    /tasks/{id}
PATCH  /tasks/{id}
DELETE /tasks/{id}
```

---

## Phase 5 — Databases

Goal: persist application data using PostgreSQL.

### Topics

- [ ] `database/sql`
- [ ] PostgreSQL connection
- [ ] Connection pooling
- [ ] Queries
- [ ] Inserts
- [ ] Updates
- [ ] Deletes
- [ ] Transactions
- [ ] Context-aware queries
- [ ] Database errors
- [ ] Repository pattern
- [ ] Database migrations
- [ ] Integration testing

### Phase 5 Project

Add PostgreSQL persistence to the task API.

Features:

- Tasks stored in PostgreSQL
- Database migrations
- Transactions
- Repository abstraction
- Integration tests
- Docker Compose environment

---

## Phase 6 — Concurrency

Goal: understand Go's concurrency model and safely run tasks in parallel.

### Topics

- [ ] Goroutines
- [ ] Channels
- [ ] Buffered channels
- [ ] Unbuffered channels
- [ ] Channel direction
- [ ] Closing channels
- [ ] Comma-ok with channels
- [ ] `select`
- [ ] `sync.WaitGroup`
- [ ] `sync.Mutex`
- [ ] `sync.RWMutex`
- [ ] Worker pools
- [ ] Fan-out
- [ ] Fan-in
- [ ] Context cancellation
- [ ] Timeouts
- [ ] Backpressure
- [ ] Race conditions
- [ ] Goroutine leaks

### Phase 6 Project

Build a concurrent URL checker:

- Read URLs from a file
- Process URLs with workers
- Limit concurrency
- Use request timeouts
- Collect response statistics
- Support cancellation
- Avoid goroutine leaks

---

## Phase 7 — Production Backend Development

Goal: build production-ready Go services.

### Topics

- [ ] Application configuration
- [ ] Environment variables
- [ ] Structured logging with `log/slog`
- [ ] Graceful shutdown
- [ ] Health endpoints
- [ ] Readiness endpoints
- [ ] Authentication
- [ ] Authorization
- [ ] Password hashing
- [ ] Retry strategies
- [ ] Idempotency
- [ ] Rate limiting
- [ ] Observability
- [ ] Metrics
- [ ] Tracing
- [ ] Security checks
- [ ] Vulnerability scanning
- [ ] Docker
- [ ] Multi-stage builds
- [ ] Deployment fundamentals

### Important Commands

```bash
go fmt ./...
go vet ./...
go test ./...
go test -race ./...
govulncheck ./...
```

---

## Final Project — Job Processing Service

Build a production-style background job system.

### API

```text
POST /jobs
GET  /jobs/{id}
POST /jobs/{id}/cancel
GET  /health
GET  /ready
```

### Requirements

- PostgreSQL persistence
- Worker pool
- Goroutines and channels
- Job retries
- Cancellation
- Context propagation
- Graceful shutdown
- Structured logs
- Unit tests
- Integration tests
- Docker Compose
- Health checks
- Metrics
- Clean project structure
- Architecture documentation

---

# Current Progress

We are currently finishing:

```text
Phase 1 — Go Fundamentals
```

The next lesson is:

```text
defer, panic, and recover
```

After the remaining Phase 1 topics, we will build the command-line task manager before moving to Phase 2.

---

# Recommended Study Rhythm

A good rhythm is:

```text
Lesson
→ Exercises
→ Code review
→ Corrections
→ Small challenge
→ Next lesson
```

Recommended study time:

```text
45–90 minutes per session
3–5 sessions per week
```

The most important rule is:

```text
Write code during every lesson.
```

Watching or reading without practicing should be avoided.

---

# Course Outcome

By the end of the course, you should be able to:

- Write idiomatic Go
- Build command-line applications
- Organize Go projects
- Write automated tests
- Build REST APIs
- Work with PostgreSQL
- Use goroutines and channels safely
- Handle contexts and cancellation
- Build production-ready backend services
- Containerize and deploy Go applications
