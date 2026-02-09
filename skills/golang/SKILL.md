---
name: golang
description: Go language development assistance. Use when working with Go code, including writing, refactoring, debugging, testing, and optimizing Go programs. Covers standard library usage, common patterns, module management, concurrency, and Go best practices.
---

# Go Development

## Core Principles

### Idiomatic Go

Write simple, clear, and idiomatic Go code:

- **Keep it simple**: Go values simplicity over cleverness
- **Explicit over implicit**: Avoid magic, be clear about what code does
- **Composition over inheritance**: Use interfaces and embedding, not inheritance
- **Error handling**: Always check errors, don't panic unless necessary

### Code Organization

```
project/
├── main.go              # Entry point (for applications)
├── go.mod               # Module definition
├── go.sum               # Dependency checksums
├── internal/            # Private application code
│   └── pkg/
├── pkg/                 # Public library code
├── cmd/                 # Application entry points
│   └── app/
│       └── main.go
└── api/                 # API definitions
```

## Common Patterns

### Error Handling

```go
// Wrap errors with context
if err != nil {
    return fmt.Errorf("failed to process %s: %w", name, err)
}

// Define sentinel errors
var ErrNotFound = errors.New("not found")

// Check error types
if errors.Is(err, ErrNotFound) {
    // handle not found
}

// Custom error types
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}
```

### Structs and Interfaces

```go
// Small, focused interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Interface satisfaction is implicit
type MyReader struct{}

func (m *MyReader) Read(p []byte) (n int, err error) {
    return 0, nil
}

// Constructor pattern
type Service struct {
    client *http.Client
    config Config
}

func NewService(cfg Config) *Service {
    return &Service{
        client: &http.Client{Timeout: cfg.Timeout},
        config: cfg,
    }
}
```

### Concurrency

```go
// Goroutines with proper cleanup
func processItems(items []Item) error {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    errChan := make(chan error, len(items))
    var wg sync.WaitGroup
    
    sem := make(chan struct{}, 10) // Limit concurrency
    
    for _, item := range items {
        wg.Add(1)
        go func(i Item) {
            defer wg.Done()
            sem <- struct{}{}        // Acquire
            defer func() { <-sem }() // Release
            
            if err := process(ctx, i); err != nil {
                errChan <- err
            }
        }(item)
    }
    
    go func() {
        wg.Wait()
        close(errChan)
    }()
    
    for err := range errChan {
        if err != nil {
            return err
        }
    }
    return nil
}
```

## Module Management

### Common Commands

```bash
# Initialize module
go mod init module-name

# Add dependency
go get github.com/pkg/example

# Update dependencies
go get -u ./...

# Tidy modules
go mod tidy

# Vendor dependencies
go mod vendor

# List modules
go list -m all
```

### Module Best Practices

- Use semantic versioning tags for releases
- Keep module paths clean and meaningful
- Minimize external dependencies
- Use `replace` directive only for local development

## Testing

### Table-Driven Tests

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -2, -3, -5},
        {"zero", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.expected {
                t.Errorf("Add(%d, %d) = %d, want %d", 
                    tt.a, tt.b, got, tt.expected)
            }
        })
    }
}
```

### Benchmarking

```go
func BenchmarkProcess(b *testing.B) {
    data := generateData()
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        Process(data)
    }
}
```

## Common Libraries

| Purpose | Library |
|---------|---------|
| HTTP Router | `github.com/go-chi/chi` |
| Configuration | `github.com/spf13/viper` |
| CLI | `github.com/spf13/cobra` |
| Logging | `log/slog` (std), `go.uber.org/zap` |
| Database | `database/sql`, `github.com/jmoiron/sqlx` |
| Migrations | `github.com/golang-migrate/migrate` |
| Validation | `github.com/go-playground/validator` |
| Testing | `github.com/stretchr/testify` |

## References

- **Effective Go**: [references/effective-go.md](references/effective-go.md)
- **Common Pitfalls**: [references/pitfalls.md](references/pitfalls.md)
- **Performance Tips**: [references/performance.md](references/performance.md)
