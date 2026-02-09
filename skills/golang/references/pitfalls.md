# Common Go Pitfalls

## Concurrency Issues

### Goroutine Leaks

```go
// BAD: Goroutine leak - no way to stop
func process() {
    ch := make(chan int)
    go func() {
        ch <- expensiveOperation()
    }()
    // If caller returns early, goroutine is leaked
}

// GOOD: Use context for cancellation
func process(ctx context.Context) error {
    ch := make(chan int, 1)
    go func() {
        select {
        case ch <- expensiveOperation():
        case <-ctx.Done():
        }
    }()
    
    select {
    case result := <-ch:
        _ = result
    case <-ctx.Done():
        return ctx.Err()
    }
    return nil
}
```

### Closing Channels

```go
// Only sender should close channel
// Closing an already closed channel panics
// Sending to a closed channel panics

// BAD
ch := make(chan int)
go func() {
    ch <- 1
}()
close(ch)  // Race condition!

// GOOD
ch := make(chan int)
go func() {
    defer close(ch)
    ch <- 1
}()
```

## Map Concurrency

```go
// Maps are not safe for concurrent use!

// BAD: Concurrent read/write
var m = make(map[string]int)

go func() { m["a"] = 1 }()
go func() { _ = m["a"] }()  // Data race!

// GOOD: Use sync.Map or mutex
var m = sync.Map{}

// Or with mutex
var (
    mu sync.RWMutex
    m  = make(map[string]int)
)

mu.Lock()
m["a"] = 1
mu.Unlock()
```

## Nil Pointers

```go
// nil interface is not nil!
var p *MyStruct = nil
var i interface{} = p

fmt.Println(i == nil)  // false!

// Solution: return nil interface directly
func GetValue() interface{} {
    var p *MyStruct = nil
    if p == nil {
        return nil  // Return untyped nil
    }
    return p
}
```

## JSON Handling

```go
// Unexported fields are ignored
type Config struct {
    APIKey string  // ignored
    apiKey string  // ignored too
}

// Use pointers for optional fields
type Request struct {
    Name  *string `json:"name,omitempty"`
    Count *int    `json:"count,omitempty"`
}

// Custom marshaling
type TimeUnix time.Time

func (t TimeUnix) MarshalJSON() ([]byte, error) {
    return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}
```

## Error Handling

```go
// Don't use panic for normal errors
// Don't ignore errors with _

// BAD
f, _ := os.Open("file")  // Silent failure

// GOOD
f, err := os.Open("file")
if err != nil {
    return err
}
defer f.Close()
```
