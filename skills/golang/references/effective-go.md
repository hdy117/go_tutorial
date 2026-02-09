# Effective Go Reference

## Formatting

- Use `gofmt` to format code
- Use `goimports` to manage imports
- Line length: aim for ~100 characters, but don't be rigid

## Naming Conventions

```go
// Public: PascalCase
func PublicFunction() {}
type PublicStruct struct{}

// Private: camelCase
func privateFunction() {}
type privateStruct struct{}

// Acronyms: all caps
HTTPRequest, URLString, IDGenerator

// Getters: use field name, not GetX
type Person struct {
    name string
}

func (p *Person) Name() string { return p.name }  // Not GetName

// Setters: use Set prefix
func (p *Person) SetName(name string) { p.name = name }
```

## Control Structures

```go
// If with short statement
if err := doSomething(); err != nil {
    return err
}

// Switch without expression (cleaner than if-else chain)
switch {
    case x < 0:
        return -1
    case x > 0:
        return 1
    default:
        return 0
}

// Type switch
switch v := i.(type) {
case int:
    fmt.Printf("int: %d\n", v)
case string:
    fmt.Printf("string: %s\n", v)
default:
    fmt.Printf("unknown: %T\n", v)
}
```

## Slices and Maps

```go
// Preallocate when size is known
s := make([]int, 0, 100)

// Check map existence with comma ok
if v, ok := m["key"]; ok {
    // use v
}

// Delete from map
delete(m, "key")

// Copy slice
dst := make([]int, len(src))
copy(dst, src)
```

## Defer

```go
// Clean up resources
f, err := os.Open("file")
if err != nil {
    return err
}
defer f.Close()

// Deferred calls run LIFO
defer fmt.Println("1")
defer fmt.Println("2")  // Prints 2, then 1

// Defer with function (evaluates args immediately)
defer func() {
    // cleanup
}()
```

## Pointers vs Values

```go
// Use pointers for:
// - Large structs
// - When mutation is needed
// - Consistency when some methods need pointers

// Use values for:
// - Small structs (<= 64 bytes)
// - Immutability
// - When nil is not desired
```
