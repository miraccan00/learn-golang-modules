# Golang Modules User Guide

This guide covers the process of creating modules, managing dependencies, and versioning modules in the Go programming language. Golang Modules make it easy to create independent packages and share code between projects.

---

## 📌 What are Golang Modules?

Golang Modules are a system used for dependency management and project versioning in the Go language. They were introduced in Go 1.11 and became the default dependency management tool from version 1.13 onwards. With modules, you can:

- Organize your projects into independent packages.
- Manage external dependencies (version control, downloading, updating).
- Facilitate code sharing between different projects.

---

## 🛠️ 1. Creating a Module: First Steps

### 📂 Folder Structure:

```plaintext
learn-golang-modules
    ├── go.mod
    ├── main.go
    ├── goodbye
    ├── greeting
    └── hello
```

In this structure, three separate modules named "goodbye", "greeting", and "hello" have been created.

### 📋 Steps:

#### A. Create the Project Folder:
```bash
mkdir learn-golang-modules
cd learn-golang-modules
```

#### B. Initialize the Module:
```bash
go mod init https://github.com/miraccan00/learn-golang-modules.git
```
- `https://github.com/miraccan00/learn-golang-modules.git`: Module name (does not have to be a real domain).

#### C. go.mod File:
Automatically created with the following content:

```go
module https://github.com/miraccan00/learn-golang-modules.git

go 1.24
```

---

## 🚀 2. Writing Your First Go Code

### Create main.go File:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go Modules!")
}
```

### Run the Application:
```bash
go run main.go
```

**Expected Output:**
```plaintext
Hello, Go Modules!
```

---

## 🔗 3. Adding an External Package
For example, let's add the `github.com/fatih/color` package.

```bash
go get github.com/fatih/color
```

### Update the Code:

```go
package main

import (
    "fmt"
    "github.com/fatih/color"
)

func main() {
    fmt.Println("Normal text")
    color.Red("This is red text")
    color.Green("And this is green!")
}
```

### Run Again:
```bash
go run main.go
```

---

## 🏷️ 4. Module Versioning

### Using Git Tags:

```bash
# Create a version
git tag v1.0.0

# Push to remote repository
git push origin v1.0.0
```

### Create a New Version on GitHub:

1. Go to the repo page on GitHub.
2. In the **Releases** section, create a new **Tag**.
3. Set the tag name to something like `v1.0.0`.

---

## 🧠 Useful Go Mod Commands

| Command | Description |
|---------|-------------|
| `go mod init <module_name>` | Initializes a new module. |
| `go mod tidy` | Cleans up unused dependencies. |
| `go get <package>` | Downloads an external package. |
| `go test ./...` | Runs tests in all packages. |

---

## How to Use This Package
```bash
go get https://github.com/miraccan00/learn-golang-modules@v1.0.0
```


## 🎯 Conclusion

With this guide, you learned how to:
- Create Go modules from scratch.
- Manage external dependencies.
- Perform version management using Git tags.

For more information, you can visit the official Go documentation: [Golang Modules](https://golang.org/doc/go1.11#modules)

