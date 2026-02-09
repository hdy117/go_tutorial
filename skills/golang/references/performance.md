# Go Performance Tips

## Memory Optimization

### Preallocation

```go
// BAD: Multiple allocations
var s []int
for i := 0; i < 1000; i++ {
    s = append(s, i)  // Reallocates multiple times
}

// GOOD: Single allocation
s := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    s = append(s, i)
}
```

### String Building

```go
// BAD: O(n²) string concatenation
var s string
for i := 0; i < 1000; i++ {
    s += "x"  // Allocates new string each time
}

// GOOD: Use strings.Builder
var b strings.Builder
b.Grow(1000)  // Preallocate
for i := 0; i < 1000; i++ {
    b.WriteString("x")
}
s := b.String()
```

### Avoiding Allocations

```go
// Escape analysis: keep values on stack
func NoEscape() {
    x := 5
    println(x)  // Stays on stack
}

func DoesEscape() *int {
    x := 5
    return &x  // Escapes to heap
}
```

## Concurrency Performance

### Worker Pools

```go
func workerPool(jobs []Job, workers int) {
    jobChan := make(chan Job, len(jobs))
    resultChan := make(chan Result, len(jobs))
    
    // Start workers
    var wg sync.WaitGroup
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobChan {
                resultChan <- process(job)
            }
        }()
    }
    
    // Send jobs
    for _, job := range jobs {
        jobChan <- job
    }
    close(jobChan)
    
    // Collect results
    go func() {
        wg.Wait()
        close(resultChan)
    }()
    
    for r := range resultChan {
        use(r)
    }
}
```

### Sync.Pool

```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 4096)
    },
}

func process() {
    buf := bufferPool.Get().([]byte)
    defer bufferPool.Put(buf)
    
    // Use buf...
}
```

## Profiling

```go
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    // ...
}
```

### Commands

```bash
# CPU profile
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# Memory profile
go tool pprof http://localhost:6060/debug/pprof/heap

# Trace
curl -o trace.out http://localhost:6060/debug/pprof/trace?seconds=5
go tool trace trace.out
```

## Benchmarking Best Practices

```go
func BenchmarkSliceAppend(b *testing.B) {
    // Reset timer after setup
    data := make([]int, 100)
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        _ = append([]int(nil), data...)
    }
}

func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            doWork()
        }
    })
}
```
