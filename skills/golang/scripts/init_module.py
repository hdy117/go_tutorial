#!/usr/bin/env python3
"""
Initialize a new Go module with standard project structure.
Usage: python3 init_module.py <module-name>
"""

import os
import sys
import subprocess


def run(cmd, **kwargs):
    """Run a shell command."""
    result = subprocess.run(cmd, shell=True, capture_output=True, text=True, **kwargs)
    if result.returncode != 0:
        print(f"Error: {result.stderr}")
        sys.exit(1)
    return result.stdout


def main():
    if len(sys.argv) < 2:
        print("Usage: python3 init_module.py <module-name>")
        sys.exit(1)

    module_name = sys.argv[1]

    # Check if go is installed
    try:
        run("go version")
    except:
        print("Error: Go is not installed or not in PATH")
        sys.exit(1)

    # Initialize module
    run(f"go mod init {module_name}")

    # Create directory structure
    dirs = ["cmd", "internal", "pkg", "api", "scripts", "configs"]
    for d in dirs:
        os.makedirs(d, exist_ok=True)

    # Create main.go template
    main_go = f'''package main

import (
	"fmt"
	"log"
)

func main() {{
	fmt.Println("Hello, {module_name}!")
	if err := run(); err != nil {{
		log.Fatal(err)
	}}
}}

func run() error {{
	// TODO: Implement
	return nil
}}
'''

    with open("main.go", "w") as f:
        f.write(main_go)

    # Create .gitignore
    gitignore = '''# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Go workspace file
go.work

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
'''

    with open(".gitignore", "w") as f:
        f.write(gitignore)

    # Create Makefile
    makefile = '''.PHONY: build test clean lint run

BINARY_NAME={module_name}
BUILD_DIR=./build

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) -v

test:
	go test -v -race ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean:
	go clean
	rm -rf $(BUILD_DIR)

lint:
	golangci-lint run

run: build
	$(BUILD_DIR)/$(BINARY_NAME)

deps:
	go mod download
	go mod tidy

.DEFAULT_GOAL := build
'''.format(module_name=module_name)

    with open("Makefile", "w") as f:
        f.write(makefile)

    # Create README.md
    readme = f'''# {module_name}

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

```bash
go mod download
```

### Build

```bash
make build
```

### Run

```bash
make run
```

### Test

```bash
make test
```

## Project Structure

```
.
├── cmd/          # Application entrypoints
├── internal/     # Private application code
├── pkg/          # Public library code
├── api/          # API definitions
├── configs/      # Configuration files
└── scripts/      # Build and utility scripts
```
'''

    with open("README.md", "w") as f:
        f.write(readme)

    # Run go mod tidy
    run("go mod tidy")

    print(f"✅ Initialized Go module: {module_name}")
    print("\nProject structure:")
    run("tree -L 1 -a || ls -la")


if __name__ == "__main__":
    main()
