# Go Course Handoff

This repository is being used as a guided, practical Go course.

The student is already a developer and is learning Go through explanation, implementation, review, correction, and incremental projects.

---

## Teaching Method

Each lesson should follow this sequence:

1. Explain one focused concept.
2. Show small, practical examples.
3. Give two or three exercises.
4. Give one output-prediction challenge.
5. Wait for the student’s attempt.
6. Review correctness first.
7. Explain why corrections are necessary.
8. Suggest idiomatic Go improvements separately.
9. Continue only after the student confirms.

Do not provide full exercise solutions before the student attempts them.

When the student asks for deeper explanation, explain the behavior at both language and memory/runtime level when relevant.

---

# Completed Course Topics

## Phase 1 — Go Fundamentals

### Program structure

Completed:

- [x] `package main`
- [x] `func main()`
- [x] Imports
- [x] Running with `go run .`
- [x] Difference between `go run .` and `go run main.go`
- [x] WSL and VS Code setup
- [x] Opening a project in a WSL-connected VS Code window
- [x] Installing Go inside WSL
- [x] Using the official Go extension in WSL

### Variables and basic types

Completed:

- [x] `var`
- [x] `:=`
- [x] Assignment with `=`
- [x] Type inference
- [x] Constants
- [x] Zero values
- [x] Multiple assignment
- [x] Basic types:
- [x] `string`
- [x] `int`
- [x] `float64`
- [x] `bool`
- [x] `rune`
- [x] `byte`

### Functions

Completed:

- [x] Function parameters
- [x] Return values
- [x] Multiple return values
- [x] Returning expressions directly
- [x] Matching expression types to declared return types
- [x] Boolean expressions such as:

```go
return product.Stock > 0
```

Important understanding:

- Expressions are always evaluated.
- The operator determines the result type.
- The result type must match the function return type.

### Control flow

Completed:

- [x] Comparison operators
- [x] `if`
- [x] `else`
- [x] `else if`
- [x] Logical operators:
- [x] `&&`
- [x] `||`
- [x] `!`
- [x] Short declarations inside `if`
- [x] `switch`
- [x] Value-less `switch`
- [x] Case ordering

### Loops

Completed:

- [x] Standard `for`
- [x] While-style `for`
- [x] Infinite loops
- [x] `break`
- [x] `continue`
- [x] Accumulation
- [x] FizzBuzz
- [x] Counting even numbers

### Arrays and slices

Completed:

- [x] Arrays
- [x] Slice literals
- [x] `append`
- [x] `make`
- [x] `len`
- [x] `cap`
- [x] Indexing
- [x] `range`
- [x] Updating by index
- [x] Slice expressions
- [x] Shared backing arrays
- [x] Copying slices with `copy`

Important understanding:

- A slice is a descriptor over an underlying array.
- Reslicing does not normally copy all values.
- Two slices can reference the same backing array.
- Updating one can affect the other.

### Maps

Completed:

- [x] Map literals
- [x] `make(map[...])`
- [x] Add and update
- [x] Delete
- [x] Iteration
- [x] Missing keys and zero values
- [x] Maps as reference-like values
- [x] Maps of slices
- [x] Map key restrictions

### Comma-ok idiom

Completed:

- [x] `value, ok := myMap[key]`

Important understanding:

- For maps, `ok` means the key exists.
- For type assertions, `ok` means the assertion succeeded.
- For channel receives, `ok` indicates whether the channel is still open and yielded a value.
- Missing map keys return the zero value plus `false`.

### Structs and methods

Completed:

- [x] Struct declarations
- [x] Named-field initialization
- [x] Zero values
- [x] Passing structs by value
- [x] Pointers to structs
- [x] Value receivers
- [x] Pointer receivers
- [x] Constructor-style functions such as `New`
- [x] Nested structs

Important understanding:

- Value receiver methods work on copies.
- Pointer receiver methods can mutate the original value.
- Go automatically handles `&value` and dereferencing in many method calls.

### Interfaces

Completed:

- [x] Interface definitions
- [x] Implicit implementation
- [x] Small interfaces
- [x] Multiple concrete implementations
- [x] Interface method sets
- [x] Value receiver vs pointer receiver behavior
- [x] `any`
- [x] Type assertions
- [x] Type switches
- [x] Interfaces near the consumer

Important understanding:

- A type satisfies an interface automatically by having the required methods.
- If an interface requires a pointer-receiver method, only the pointer type satisfies it.
- A value-receiver method belongs to both the value and pointer method sets.

### Error handling

Completed:

- [x] Returning `error`
- [x] `errors.New`
- [x] `fmt.Errorf`
- [x] Early returns
- [x] Propagating errors
- [x] Wrapping with `%w`
- [x] Sentinel errors
- [x] `errors.Is`
- [x] Returning zero values when an error occurs

### `defer`, `panic`, and `recover`

Completed:

- [x] Deferred execution
- [x] LIFO order
- [x] Cleanup
- [x] Deferred arguments evaluated immediately
- [x] Deferred closures reading variables later
- [x] Panic behavior
- [x] Deferred functions during panic
- [x] Recover inside deferred functions
- [x] Appropriate use of `panic`

Important understanding:

```go
defer fmt.Println(value)
```

captures the argument immediately.

```go
defer func() {
    fmt.Println(value)
}()
```

reads the variable when the deferred closure executes.

### Modules and packages

Completed:

- [x] `go mod init`
- [x] `go.mod`
- [x] Module path
- [x] Import path resolution
- [x] Module vs package
- [x] Package as a compile-time namespace and compilation unit
- [x] Package-level initialization
- [x] Package-level variables
- [x] `package main`
- [x] Exported and unexported identifiers
- [x] Files in the same directory compiling as one package
- [x] Package compilation and caching
- [x] Linking packages into an executable
- [x] `cmd` directories
- [x] `internal` directories

Important understanding:

- A module can contain many packages.
- A package is usually one directory of Go files.
- Files in the same package compile together.
- Packages are not runtime objects.
- Package functions become machine code.
- Package-level variables occupy runtime memory.
- Lowercase visibility is enforced at compile time.

### Strings, bytes, runes, and UTF-8

Completed:

- [x] Strings are immutable.
- [x] String indexing returns bytes.
- [x] `len(string)` returns bytes.
- [x] `[]rune(string)` counts Unicode code points.
- [x] `range` decodes UTF-8.
- [x] Byte indexes from `range`
- [x] `unicode/utf8`
- [x] `strings.Builder`
- [x] String normalization
- [x] Unicode-safe reversal

Important understanding of:

```go
text = text[:len(text)-size]
```

- `DecodeLastRuneInString` finds the last rune and its byte size.
- `WriteRune` copies that rune into a builder.
- The slice expression shortens the source string by the rune’s byte size.
- The builder and source string are separate values.
- String slicing normally creates a new string header over the same bytes rather than copying all remaining content.
- Preallocating with `builder.Grow(len(text))` reduces reallocations.

### File handling

Completed:

- [x] `os.WriteFile`
- [x] `os.ReadFile`
- [x] `os.Create`
- [x] `os.Open`
- [x] `os.OpenFile`
- [x] Append flags
- [x] File permissions
- [x] `defer file.Close()`
- [x] `bufio.Scanner`
- [x] Scanner error checking
- [x] `bufio.Writer`
- [x] `Flush`
- [x] Relative paths
- [x] Empty files vs missing files

Important edge case:

- Missing file can be handled with `os.IsNotExist`.
- Empty file is successfully read but causes JSON decoding to fail unless handled separately.

### JSON

Completed:

- [x] `json.Marshal`
- [x] `json.MarshalIndent`
- [x] `json.Unmarshal`
- [x] Struct tags
- [x] `omitempty`
- [x] `json:"-"`
- [x] Exported fields
- [x] Slices of structs
- [x] File persistence
- [x] `json.NewEncoder`
- [x] `json.NewDecoder`

### Phase 1 final project

Completed project:

```text
go-task-manager/
├── cmd/
│   └── task/
│       └── main.go
├── internal/
│   ├── storage/
│   │   └── json.go
│   └── task/
│       ├── repository.go
│       ├── service.go
│       └── task.go
├── go.mod
└── tasks.json
```

CLI commands:

```bash
go run ./cmd/task add "Task title"
go run ./cmd/task list
go run ./cmd/task complete <id>
go run ./cmd/task delete <id>
```

Implemented features:

- Add task
- List tasks
- Complete task
- Delete task
- JSON persistence
- Auto-incremented IDs
- Validation for missing and empty titles
- Validation for nonnumeric and negative IDs
- Task-not-found handling
- Unknown command handling
- Usage output
- Package separation
- `go fmt ./...`
- `go vet ./...`

---

# Current Architecture

The current dependency flow is:

```text
CLI
→ task.Service
→ task.Repository interface
→ storage.JSONRepository
→ tasks.json
```

The repository interface is defined near the consumer:

```go
package task

type Repository interface {
    Load() ([]Task, error)
    Save([]Task) error
}
```

The JSON repository stores the filename internally:

```go
type JSONRepository struct {
    filename string
}
```

Its methods must match the interface exactly:

```go
Load() ([]task.Task, error)
Save([]task.Task) error
```

The service currently provides:

```go
Add(title string) (Task, error)
List() ([]Task, error)
Complete(id int) error
Delete(id int) error
```

Sentinel errors:

```go
ErrEmptyTitle
ErrTaskNotFound
```

A compile-time interface assertion may be used:

```go
var _ task.Repository = (*JSONRepository)(nil)
```

---

# Current Course Position

Current phase:

```text
Phase 2 — Idiomatic Go and Project Structure
```

Most recently completed topic:

```text
Dependency injection and repository interfaces
```

The project has already been refactored so the CLI injects `JSONRepository` into `task.Service`.

---

# Remaining Course Roadmap

## Phase 2 — Idiomatic Go and Project Structure

Remaining topics:

- [x] Custom error types
- [x] Error design and error boundaries
- [x] Configuration
- [x] Environment variables
- [x] Constructor design
- [x] Package API design
- [ ] Interface placement
- [ ] Dependency direction
- [ ] Avoiding unnecessary interfaces
- [ ] Naming conventions
- [ ] `internal` package usage
- [ ] `cmd` package usage
- [ ] `go.mod` and `go.sum` in depth
- [ ] Dependency management
- [ ] `go fmt`
- [ ] `go vet`
- [ ] Static analysis
- [ ] Basic linting
- [ ] Refactoring service and repository boundaries
- [ ] In-memory repository implementation

Expected Phase 2 project outcome:

- A cleaner task manager architecture
- Swappable JSON and in-memory repositories
- Better error design
- Configuration through environment variables
- Clear package boundaries

## Phase 3 — Testing

Topics:

- [ ] `testing` package
- [ ] Unit tests
- [ ] Table-driven tests
- [ ] Subtests
- [ ] Test helpers
- [ ] Testing sentinel errors
- [ ] Testing wrapped errors
- [ ] Fakes and stubs
- [ ] In-memory repositories
- [ ] Testing service logic
- [ ] Testing file repositories
- [ ] Coverage
- [ ] Benchmarks
- [ ] Race detector
- [ ] Fuzz testing

Commands:

```bash
go test ./...
go test -v ./...
go test -cover ./...
go test -race ./...
go test -bench=. ./...
go test -fuzz=Fuzz ./...
```

Project outcome:

- Full tests for task service
- Repository tests
- Validation tests
- Error-path tests
- Fuzz tests

## Phase 4 — HTTP APIs

Topics:

- [ ] `net/http`
- [ ] `http.Handler`
- [ ] `http.HandlerFunc`
- [ ] Routing
- [ ] Methods
- [ ] Headers
- [ ] Request bodies
- [ ] JSON requests
- [ ] JSON responses
- [ ] Status codes
- [ ] Middleware
- [ ] API validation
- [ ] Context
- [ ] Timeouts
- [ ] Graceful shutdown
- [ ] Structured logging
- [ ] API error responses

Project outcome:

Convert the task manager into a REST API:

```text
POST   /tasks
GET    /tasks
GET    /tasks/{id}
PATCH  /tasks/{id}
DELETE /tasks/{id}
```

## Phase 5 — Databases

Topics:

- [ ] `database/sql`
- [ ] PostgreSQL
- [ ] Connection pooling
- [ ] Queries
- [ ] Inserts
- [ ] Updates
- [ ] Deletes
- [ ] Transactions
- [ ] Context-aware queries
- [ ] Database errors
- [ ] Repository pattern
- [ ] Migrations
- [ ] Integration tests
- [ ] Docker Compose

Project outcome:

- PostgreSQL-backed task API
- Migrations
- Transactions
- Integration tests

## Phase 6 — Concurrency

Topics:

- [ ] Goroutines
- [ ] Channels
- [ ] Buffered and unbuffered channels
- [ ] Directional channels
- [ ] Closing channels
- [ ] Comma-ok with channels
- [ ] `select`
- [ ] `sync.WaitGroup`
- [ ] `sync.Mutex`
- [ ] `sync.RWMutex`
- [ ] Worker pools
- [ ] Fan-out and fan-in
- [ ] Context cancellation
- [ ] Timeouts
- [ ] Backpressure
- [ ] Race conditions
- [ ] Goroutine leaks

Project outcome:

Build a concurrent URL checker with:

- Worker pool
- Request timeouts
- Rate limiting
- Cancellation
- Statistics
- Safe shutdown

## Phase 7 — Production Backend Development

Topics:

- [ ] Configuration
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
- [ ] Metrics
- [ ] Tracing
- [ ] Security checks
- [ ] Vulnerability scanning
- [ ] Docker
- [ ] Multi-stage builds
- [ ] Deployment fundamentals

Important commands:

```bash
go fmt ./...
go vet ./...
go test ./...
go test -race ./...
govulncheck ./...
```

## Final Project — Job Processing Service

API:

```text
POST /jobs
GET  /jobs/{id}
POST /jobs/{id}/cancel
GET  /health
GET  /ready
```

Requirements:

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

# Review Rules

When the student submits code:

1. Inspect the actual repository files.
2. Run `go fmt ./...`.
3. Run `go vet ./...`.
4. Run relevant tests.
5. Run relevant CLI commands.
6. Identify correctness problems first.
7. Explain the cause of each problem.
8. Distinguish required fixes from optional idiomatic improvements.
9. Avoid rewriting the entire project unless necessary.
10. Preserve the learning process.
11. Do not provide the next full exercise solution before the student attempts it.
12. Continue only after the student confirms readiness.

---
