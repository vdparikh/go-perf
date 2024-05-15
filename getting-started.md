# Getting started with GO

Go, also known as Golang, has gained significant popularity since its inception due to its simplicity, performance, and robust ecosystem. Here are several reasons why Go is a compelling choice for many programming tasks, including network programming and systems development:

Here are few reasons why GO is a good choice 
1. Simplicity and Readability
2. Built-in Concurrency
3. Fast Compilation
4. Powerful Standard Library
5. Static Typing and Efficiency - Despite its simplicity, Go does not sacrifice performance
6. Cross-Platform Support
7. Robust Tooling and Documentation
8. Strong Community and Corporate Support

While other languages like Rust also offer impressive features, particularly around memory safety, Go strikes an optimal balance between performance, ease of development, and readability. This makes it particularly well-suited for developing secure applications quickly and efficiently.

Go is a very popular choice in open source community and all this platforms were written in go
- Docker
- Kubernetes
- Prometheus
- Terraform
- etcd
- Vault
- Grafana
- InfluxDB
- CockroachDB
- Gitea
- And many more... 

## Fun Facts
- **The Go Gopher**: Designed by Renee French, the Go gopher is inspired by the mascot style of the WFMU-FM radio station. This character has charmed the Go community with its simplicity and has been adapted into various versions to represent Go at conferences and meetups globally.
- **Naming of Go**: Initially criticized for its generic nature which made internet searches difficult, the name "Go" was chosen for its simplicity and action-oriented nature, reflecting the language's efficiency. It’s a playful nod to the possibility of abandoning the project with the phrase, "then we can just let it go."

## Introduction to Go
**What is Go?**
- Go is a statically typed, compiled programming language developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.
- It is renowned for its simplicity, efficiency, and excellent support for concurrency.

**Key Features of Go:**
- Clean and readable syntax.
- Built-in support for concurrent programming with goroutines and channels.
- Robust standard library.

**Why Use Go?**
- Ideal for building scalable, high-performance web applications and microservices.
- Strongly supported by the developer community and used extensively in cloud services and infrastructure projects.


## Setting Up the Go Development Environment
**Installing Go:**
- Provides step-by-step instructions for downloading and installing Go from the official website golang.org.
- Explanation on how to set environment variables such as GOPATH, GOROOT, and GOBIN.

**Tools You’ll Need:**
- Recommendations for a text editor or IDE (e.g., Visual Studio Code with the Go extension, GoLand).
- Introduction to Go command-line tools (go build, go run, etc.).


## Creating Your First Go Program
Starting your journey with Go involves setting up your development environment and writing a simple program. Lets create your first Go program, which will print "Hello, World!" to the console.

### Setting Up Your Go Environment
Before writing your first program, ensure that Go is installed on your system. You can download and install Go from [the official Go website](https://golang.org/dl/). Follow the installation instructions specific to your operating system.

- **Verify Installation**: After installing, you can verify the installation by opening a command prompt or terminal and running:
  ```bash
  go version
  ```
  This command will print the installed version of Go.

### Writing the Program
1. **Create a New Directory for Your Go Project**:
   - Open your terminal or command prompt.
   - Create a new directory for your project and navigate into it:
     ```bash
     mkdir my-go-project
     cd my-go-project
     ```

2. **Create a New Go File**:
   - Within your project directory, create a new file called `main.go`.
   - Open this file in a text editor or IDE of your choice.

3. **Add the Following Code to `main.go`**:
   - This code defines a package named `main`, imports the `fmt` library (used for formatted I/O operations), and defines the `main` function, which is the entry point of every Go program.
   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World!")
   }
   ```

### Running Your Program
After writing the program, you can run it directly using the Go tool:

- **Run the Program**:
  - Open your terminal or command prompt.
  - Navigate to the directory containing your `main.go` file.
  - Execute the following command:
    ```bash
    go run main.go
    ```
  - This command compiles and runs your Go program in one step.

### Understanding the Code

- **Package Declaration**: Every Go file starts with a package declaration, which defines the namespace in which the file belongs. The `main` package is special, as it defines a standalone executable (not a library).
- **Import Statements**: `import "fmt"` tells Go to include the `fmt` package, which contains functions for formatted I/O, such as `Println`.
- **The `main` Function**: The `main` function is where execution of the program begins. Every standalone Go program must have a `main` function, which serves as the entry point.

### What's Next?
- **Experiment**: Try modifying the message in the `Println` function to see different outputs.
- **Learn More**: Explore more about Go's standard library, which offers various packages for different needs like handling I/O, manipulating strings, and performing network operations.



## Data Types, Variables, and Constants
Understanding data types, variables, and constants is crucial for programming in Go. 

### Data Types in Go
Go is a statically typed language, which means the type of a variable is known at compile time. Here are some of the basic data types in Go:

- **Integers**: `int`, `int8`, `int16`, `int32`, `int64`
- **Unsigned Integers**: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **Floating Point Numbers**: `float32`, `float64`
- **Boolean**: `bool`
- **Strings**: `string`
- **Complex Numbers**: `complex64`, `complex128`

### Variables
Variables are used to store data of a specific type. In Go, variables can be declared in several ways.

**Using the `var` Keyword**:

```go
var name string = "Alice"
var age int = 30
```

You can also declare multiple variables at once:

```go
var (
    firstName = "John"
    lastName  = "Doe"
    age       = 32
)
```

**Short Variable Declaration**:
This is a shorthand for declaring and initializing a variable. It uses the `:=` syntax and is often used inside functions.

```go
name := "Alice"
age := 30
```

### Constants
Constants are variables whose values cannot be changed after they are set. They are declared like variables but with the `const` keyword.

```go
const Pi = 3.14159
const (
    StatusOK       = 200
    StatusNotFound = 404
)
```

Constants can be character, string, boolean, or numeric values. They provide a way to use meaningful names for fixed values in your code, enhancing its readability and maintainability.

### Type Inference

When declaring a variable without specifying an explicit type (either by using the `:=` syntax or `var` without a type), the variable's type is inferred from the value on the right-hand side.

```go
var i = 42           // int
var f = 3.142        // float64
var g = 0.867 + 0.5i // complex128
```

### Example: Using Data Types, Variables, and Constants

```go
package main

import "fmt"

func main() {
    const Pi = 3.14159
    var radius = 2.0
    var circumference = 2 * Pi * radius

    fmt.Println("Circumference:", circumference)
}
```

### Best Practices
- **Use Short Variable Declarations**: When you need to declare and initialize a variable inside a function, consider using the `:=` syntax for brevity and readability.
- **Use Constants for Fixed Values**: Define constants for values that never change to prevent accidental modification and improve code readability.
- **Choose the Right Data Type**: Be specific about the data types you choose to optimize the performance and clarity of your code.


## Control Structures
  - `if`, `else`, and `switch` for conditional operations.
  - `for` loop — the only looping construct in Go.

```go
package main

import "fmt"

func main() {
    // If statement
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    // For loop
    for i := 0; i <= 5; i++ {
        fmt.Println("Counter:", i)
    }

    // Switch statement
    switch day := 3; day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Other day")
    }
}
```

- **Functions:**
  - Defining functions with parameters and return types.
  - Multiple return values and named return values.

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0.0 {
        return 0.0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10.0, 0.0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

## Complex Types and Composite Data Structures
Go offers several complex types that allow you to structure your data more elaborately. Understanding these is key to building more sophisticated applications.

### Arrays

Arrays in Go are fixed-size, ordered collections of elements of the same type. They are useful when you know the exact number of elements required, and they provide fast access to their elements.

**Example of an Array:**

```go
package main

import "fmt"

func main() {
    var numbers [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array:", numbers)
}
```

### Slices

Slices are more flexible, dynamic sequences that are built on top of arrays. They can grow and shrink, which makes them more versatile than arrays.

**Creating and Manipulating a Slice:**

```go
package main

import "fmt"

func main() {
    colors := []string{"red", "green", "blue"}
    fmt.Println("Initial slice:", colors)

    // Append to a slice
    colors = append(colors, "yellow")
    fmt.Println("Updated slice:", colors)

    // Slicing a slice
    subColors := colors[1:3]
    fmt.Println("Sub-slice:", subColors)
}
```

### Maps

Maps are key-value pairs that provide a powerful way to organize data dynamically. Maps are particularly useful when you need fast lookups, deletions, and updates by keys.

**Example of Using a Map:**

```go
package main

import "fmt"

func main() {
    capitals := map[string]string{
        "France": "Paris",
        "Italy": "Rome",
        "Japan": "Tokyo",
    }

    // Adding an item to the map
    capitals["India"] = "New Delhi"

    // Deleting an item
    delete(capitals, "Japan")

    // Accessing and printing the map
    for country, capital := range capitals {
        fmt.Printf("The capital of %s is %s.\n", country, capital)
    }
}
```

### Structs

Structs are a way to group related data items together, forming a more complex data item. They are useful for modeling real-world entities within programs.

**Defining and Using Structs:**

```go
package main

import "fmt"

type Person struct {
    Name    string
    Age     int
    Country string
}

func main() {
    // Creating an instance of a struct
    bob := Person{Name: "Bob", Age: 25, Country: "USA"}
    fmt.Println(bob)

    // Accessing struct fields
    fmt.Println("Name:", bob.Name)
    fmt.Println("Age:", bob.Age)

    // Updating a field
    bob.Age = 26
    fmt.Println("Updated Age:", bob.Age)
}
```

### Methods on Structs

Methods can be defined on structs. This allows you to associate functions with particular types, which can manipulate the struct's fields.

**Example of a Method on a Struct:**

```go
package main

import "fmt"

type Rectangle struct {
    Width, Height float64
}

// A method to calculate area of the rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    area := rect.Area()
    fmt.Println("Area of rectangle:", area)
}
```

### Why Learn These Types?
Understanding and using these complex and composite data structures allow you to:
- Organize data logically and efficiently.
- Implement robust and maintainable code architectures.
- Utilize Go's type system to its full extent, leading to safer and more reliable code.


## Error Handling, Logging, and Panic in Go

### Error Handling in Go
Go handles errors explicitly using the `error` type, which is a built-in interface. This approach encourages you to check for errors where they occur and handle them appropriately.

**Example of Error Handling:**

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (result float64, err error) {
    if b == 0.0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(4.0, 2.0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

### Logging in Go

Logging in Go can be handled using the `log` package, which provides a simple logging interface. You can log messages that include the date and time of each logged message.

**Example of Logging:**

```go
package main

import (
    "log"
)

func main() {
    log.Println("Starting the application...")
    log.Println("Something happened here")
    log.Println("Finishing the application...")
}
```

With Go 1.21, Structured logging comes natively without using external library. 
log/slog


```go
package main

import (
    "log"
    "log/slog"
)

func main() {
    // INFO level by default but you can customize it for lower level messages
    // You can use AddSource: true in handler options and that will add the function, file and line from where the log was generated
    opts := &slog.HandlerOptions{
        Level: slog.LevelDebug,
        AddSource: true,
    }

    

    logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
    // logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

    logger.Debug("Debug message")
    logger.Info("Info message")
    logger.Warn("Warning message")
    logger.Error("Error message")

    // You can also customize the default Logger is to utilize the slog.SetDefault() method, allowing you to replace the default logger with a custom one.
    slog.SetDefault(logger)
    slog.Info("Info message")
    log.Println("Hello from Log")

    // You can also provide additional context about the logged event, which can be valuable for tasks such as troubleshooting, generating metrics, auditing, and various other purposes
    logger.Info(
        "some request",
        "method", "GET",
        "time_taken_ms", 100,
        slog.Int("time_taken_ms", 158),
        "path", "/encrypt",
        "status", 200,
        "user_agent", "UA_AGENT",
    )
}

```

### Panic and Recover

`panic` and `recover` are built-in functions that handle unexpected errors in Go. `panic` stops the ordinary flow of execution and begins panicking, whereas `recover` can regain control of a panicking goroutine.

**Panic Example:**

```go
package main

import (
    "log"
)

func riskyFunction() {
    defer func() {
        if r := recover(); r != nil {
            log.Println("Recovered in riskyFunction:", r)
        }
    }()
    panic("a problem")
}

func main() {
    riskyFunction()
    log.Println("Returned normally from riskyFunction.")
}
```

**Explanation:**
- **Panic**: Used within `riskyFunction` to simulate an error.
- **Recover**: Placed inside a deferred function to catch the panic, log the issue, and allow the program to continue.

### Best Practices for Logging and Error Handling

- **Use standard error checking where possible**. Reserve panic for truly exceptional situations where continuation is impossible.
- **Always defer recover in the same function as the panic**. Recovery behavior won't work if you defer recovery in a separate function.
- **Use logging judiciously**. Over-logging can fill up log files and make it difficult to find important information. Under-logging can make debugging difficult. LogLevel can help in categorizing the logs.
- **Structure your logs**. This can be especially helpful in larger applications or microservices where you need to track issues across different services and requests.


## Understanding Concurrency in Go
- **Goroutines:**
  - What are goroutines and how do they enable concurrency?
  - Creating goroutines to perform concurrent tasks.
```go
package main

import (
    "fmt"
    "time"
)

func printCounts(label string) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("%s: %d\n", label, i)
        time.Sleep(time.Millisecond * 500) // Simulate work
    }
}

func main() {
    go printCounts("goroutine1")
    go printCounts("goroutine2")

    // Wait for input to give goroutines time to finish
    fmt.Scanln()
    fmt.Println("Main function finished")
}
```
- **Channels:**
  - Using channels to communicate between goroutines.
  - Buffered and unbuffered channels.

```go
package main

import "fmt"

func sendMessage(ch chan string) {
    ch <- "Hello from goroutine!"
}

func main() {
    messageChan := make(chan string)
    go sendMessage(messageChan)

    // Receive message from the channel
    message := <-messageChan
    fmt.Println(message)
}
```

- **Synchronization Primitives:**
  - Using `sync.Mutex` and `sync.WaitGroup` to synchronize access to shared resources.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers completed")
}
```

## Building a Simple Project
- **Project Idea:**
  - Developing a simple HTTP server.
  - Handling basic web routes.

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

- **Testing the Application:**
  - How to access the server via a web browser.
    - Once the server is running, open any web browser.
    - Type in the URL http://localhost:8080 in the address bar (assuming the server is running on port 8080).
    - If everything is set up correctly, the browser should display the response defined in your server's handler function, e.g., "Hello, World!" or any other output specified.
  - Tips on debugging common issues.
    - Add Logging statements
        ```go
        func handler(w http.ResponseWriter, r *http.Request) {
            log.Println("Received request from", r.RemoteAddr)
            fmt.Fprintf(w, "Hello, World!")
            log.Println("Handled and responded to request successfully")
        }
        ```


## Putting your application on Docker
Docker is a platform for developing, shipping, and running applications inside lightweight, portable containers.

### What is Dockerfile?
A Dockerfile is a text document that contains all the commands a user could call on the command line to assemble an image.

- Create a dockerfile
    ```dockerfile
    # Start from the official Go image.
    FROM golang:1.16

    # Set the Current Working Directory inside the container
    WORKDIR /app

    # Copy everything from the current directory to the Working Directory inside the container
    COPY . .

    # Download all the dependencies that are necessary
    RUN go mod download

    # Build the Go app
    RUN go build -o main .

    # Expose port 8080 to the outside world
    EXPOSE 8080

    # Command to run the executable
    CMD ["./main"]
    ```

- Build your docker image `docker build -t my-go-app .`
- Run your image `docker run -p 8080:8080 my-go-app`


## Unit Testing and Test-Driven Development (TDD) in Go
Go has a powerful native testing framework, which includes a package called `testing`. This allows you to easily write automated tests for your code.

**Example of a Simple Test:**

```go
package main

import (
    "testing"
    "strings"
)

func TestToUpper(t *testing.T) {
    result := strings.ToUpper("test")
    if result != "TEST" {
        t.Errorf("Expected TEST, got %s", result)
    }
}
```

You can run tests by executing `go test` in your command line. This will automatically recognize any file ending in _test.go and execute the appropriate tests.

Benefits of Testing:
- Ensures your code works as expected.
- Makes refactoring and maintaining code safer and easier.
- Helps document the expected behavior of your functions.

## Go Modules and Package Management
Since Go modules were introduced in Go 1.11, they have become the standard for dependency management in Go, enabling reliable versioning and package management.

**Go Modules and Dependency Management**

Go Modules help manage dependencies in Go projects. A module is a collection of Go packages stored in a file tree with a `go.mod` file at its root.

**Initializing a New Module:**

```bash
go mod init # init
go mod tidy #pulls packages
```

**Why Use Go Modules?**
- Reproducible builds: Ensures your project builds with specific versions of dependencies.
- Improved dependency resolution: Automatically finds compatible versions of dependencies.
- Eliminates the need for GOPATH: Allows you to work outside of the GOPATH and store projects anywhere on your filesystem.

## Effective Go Practices

Writing idiomatic Go code means following the conventions and practices that the Go community has developed over time.

**Key Practices Include:**

- **Formatting**: Use `go fmt` to automatically format your code according to the Go standards.
- **Naming Conventions**: Follow Go's naming conventions, e.g., use MixedCaps for function names and variable names.
- **Error Handling**: Prefer returning errors as opposed to using exceptions for normal error handling.
- **Concurrency**: Utilize Go's concurrency features like goroutines and channels to write efficient concurrent code without engaging in common threading issues.


## Go Environment and Tools

Understanding the Go environment and leveraging Go's rich toolset can greatly enhance your productivity and code quality.

**Key Tools Include:**

- **Go Vet**: Analyzes your code to detect suspicious constructs that could potentially be bugs.
- **Go Lint**: Offers style suggestions to help you write cleaner and more idiomatic Go code.
- **Go Doc**: Helps you generate documentation for your packages based on comments in your code.


## Do's and Don'ts of Go Programming

Adhering to certain do's and don'ts while programming in Go can help you avoid common mistakes and embrace best practices. Here's a list of essential tips to help you along your Go programming journey.

### Do's

1. **Do Use `gofmt` to Format Your Code**:
   - Always run your code through `gofmt` or configure your IDE to do this automatically. This tool formats Go code according to the official Go style guidelines, ensuring consistency across your codebase.

   ```bash
   gofmt -w yourfile.go
   ```

2. **Do Handle Errors Where They Occur**:
   - Unlike many other languages that use exceptions, Go handles errors explicitly. Check errors where they occur, and handle them appropriately.

   ```go
   val, err := someFunction()
   if err != nil {
       // Handle the error appropriately
       log.Fatal(err)
   }
   ```

3. **Do Use Go Routines for Concurrency**:
   - Make use of Go's built-in support for concurrency with goroutines. They are lightweight and efficient, making them ideal for developing concurrent applications.

   ```go
   go func() {
       fmt.Println("I'm running in parallel!")
   }()
   ```

4. **Do Use Proper Naming Conventions**:
   - Follow Go's naming conventions: PascalCase for exported functions and variables, camelCase for internal functions and variables.

5. **Do Utilize Go Modules for Dependency Management**:
   - Use Go modules to manage dependencies. It’s effective and the officially recommended way to handle packages and versions.

   ```bash
   go mod init mymodule
   ```

### Don'ts

1. **Don't Ignore Go Routines' Errors**:
   - When using goroutines, make sure you have a strategy to handle errors, such as using channels or the `errgroup` package.

   ```go
   var g errgroup.Group
   g.Go(func() error {
       return errors.New("something went wrong")
   })
   if err := g.Wait(); err != nil {
       log.Printf("Encountered an error: %v", err)
   }
   ```

2. **Don't Use Panic and Recover as a Substitute for Error Handling**:
   - Avoid using `panic` and `recover` except in truly exceptional situations. Regular error handling should be done with Go’s error type.

3. **Don't Ignore Concurrency Issues**:
   - Be cautious of race conditions and data races when using goroutines and channels. Use synchronization tools like mutexes and wait groups wisely.

   ```go
   var mu sync.Mutex
   mu.Lock()
   sharedResource := someValue
   mu.Unlock()
   ```

4. **Don't Rely on Garbage Collection for Resource Management**:
   - While Go has a garbage collector, you should still manage resources explicitly, like closing files or network connections after use.

   ```go
   f, err := os.Open("filename")
   if err != nil {
       log.Fatal(err)
   }
   defer f.Close() // Ensures the file is closed
   ```

5. **Don't Overlook the Standard Library**:
   - Make use of Go's extensive standard library. Often, there's no need to reinvent the wheel with third-party packages.


## Further Learning
- [Go Playground](https://play.golang.org/) is a web service that lets your type and run Go code online. Useful for sharing snippets or experimenting with Go syntax without any setup.
- Recommended books, online courses, and tutorials.
  - golang.org/doc
  - Go by Example: Hands-on introduction to Go using annotated example programs.
  - Tour of Go: An interactive tour that walks you through the fundamentals of the language.
  - Effective Go: Provides tips on writing clear, idiomatic Go code.
- [Awesome Go](https://github.com/avelino/awesome-go)