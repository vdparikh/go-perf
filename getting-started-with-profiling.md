# Getting Started with Profiling and Go
This doc will walk you through basic profiling in Go

The full code can be found in `cmd/profiling/main.go`

https://github.com/google/pprof/tree/main

Profiling Go application helps you find hot paths. 
- Which functions are consuming most CPU and where it spends most time.
- Memory consumption
  -  The number of values that we’re throwing in the heap (number of objects, how short lived they are) 
  -  Overall heap size (how much memory do we actually need to store this amount of data).

pprof lets you collect CPU profiles, traces, and heap profiles for your Go programs. All you need to do is include `net/http/pprof` package and expose pprof endpoints.

### Prerequisites
- You will need graphviz if you want to have visualization and web output for the pprof command `brew install graphviz`. If you don't have xcode utils setup, it may also ask you to do `xcode-select --install` for graphviz. 

### Profiles
pprof exposes few different profiles. A Profile is a collection of stack traces showing the call sequences that led to instances of a particular event, such as allocation. 

- CPU: profile?seconds=10
- Memory: heap
- Goroutines: goroutine
- Goroutine blocking: block
- Locks: mutex
- Tracing: trace?seconds=5

### Running and Profiling the Application
Start your application by running go run main.go and **simulate load** using the load testing tool `hey -n 1000 -c 100 -z 30s http://localhost:8080/encrypt`. 

Run the command to fetch the CPU profile
#### CPU Profile:
```bash
go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
```

This command fetches a 30-second CPU profile. You can analyze it using interactive commands within the pprof tool.

- Top: This command shows the functions that consume the most CPU time.
- Graph: Provides a visual representation of the call graph, highlighting functions that take the most time.
- Peek: Allows you to look at specific functions and see what proportion of time they are responsible for.
- List: Shows the source code of specific functions annotated with the number of times they have been executed and the time spent in each part. (ex: `list encryptData`)

If you want to open a web browser with charts, you can use the `-http=:8081` option which will open your browser window http://localhost:8081. You'll see the pprof web interface, which allows you to explore the CPU profile. 

#### Heap Profile:
```bash
go tool pprof http://localhost:6060/debug/pprof/heap
```
This retrieves the memory allocation profile at the time of the request.

#### Goroutine Profile:
```bash
go tool pprof http://localhost:6060/debug/pprof/goroutine
```
Shows a stack trace of all current goroutines.

#### Block Profile:
```bash
go tool pprof http://localhost:6060/debug/pprof/block
```
Fetches a profile that shows where goroutines block on synchronization primitives.

### Interpreting Results
If you are on the terminal interactive UI, you can use the same top, list commands. Using web command will open the web UI 

- CPU Profiling: Look for functions that consume more CPU time than expected. These are potential candidates for optimization.
- Memory Profiling: Identify where large allocations happen and if there are unexpected memory usages suggesting memory leaks or inefficient memory use.
- Block Profiling: Useful for diagnosing deadlocks and contention issues where goroutines wait on synchronization primitives.
- Goroutine Profiling: Helps in understanding concurrency issues by showing how many and which goroutines are in which state.

### Running with example
#### Step 1: 
Start by collecting both CPU and memory profiles while your application is under load (execute load test `hey -n 1000 -c 100 -z 40s http://localhost:8080/encrypt`). 

Execute the below command
- For CPU Profile: go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
- For Memory Profile: go tool pprof http://localhost:8080/debug/pprof/heap

#### Step 2: Analyze the Memory Profile
The memory profile is particularly important for identifying issues related to excessive memory allocation. 

Here's what you should do:

- Open the Memory Profile:
- Identify High Allocators: Use the top command to see which functions are allocating the most memory. Look for the encryptData function or any function that shows a high number of allocations or a large amount of allocated memory.
- Dive Deeper into Specific Functions: Use the list command followed by the function name, e.g., `list encryptData`. This will show you the source code of the function with the amount of memory allocated on each line.
- Look for lines where new slices are created, particularly inside loops or repeated operations.

#### Step 3: Interpret the Data
From the list encryptData output, you might see something like:

```plaintext
ROUTINE ======================== main.encryptData in /path/to/your/file.go
   10MB       50MB (flat, cum) 99.9% of Total
         .
         .
      10: encrypted := make([]byte, len(data))
         .
         .
```

- Flat: The amount of memory allocated by this line itself.
- Cum: Total memory allocated by this line and all the functions it calls.

If you notice high values on a line where a new slice is allocated inside a loop or a repeated function call, it indicates that this allocation is a prime candidate for optimization.

```bash
(pprof) list encryptData
Total: 65.08s
ROUTINE ======================== main.encryptData in /Users/vishal/go/src/github.com/vdparikh/goperf/go-perf/main.go
      10ms      7.11s (flat, cum) 10.93% of Total
         .          .     48:func encryptData(data []byte, key []byte) (string, error) {
      10ms      140ms     49:	block, err := aes.NewCipher(key)
         .          .     50:	if err != nil {
         .          .     51:		return "", err
         .          .     52:	}
         .       60ms     53:	ciphertext := make([]byte, aes.BlockSize+len(data))
         .          .     54:	iv := ciphertext[:aes.BlockSize]
         .      6.79s     55:	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
         .          .     56:		return "", err
         .          .     57:	}
         .          .     58:
         .       40ms     59:	stream := cipher.NewCFBEncrypter(block, iv)
         .          .     60:	// +++++ PROFILING +++++
         .          .     61:	// Intentionally inefficient: allocate new slice each iteration
         .       20ms     62:	encrypted := make([]byte, len(data))
         .       30ms     63:	stream.XORKeyStream(encrypted, data)
         .          .     64:	ciphertext = append(iv, encrypted...)
         .          .     65:
         .          .     66:	// More efficient: use the existing slice
         .          .     67:	// stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
         .       30ms     68:	return hex.EncodeToString(ciphertext), nil
         .          .     69:
         .          .     70:}
         .          .     71:
         .          .     72:func encryptHandler(w http.ResponseWriter, r *http.Request) {
         .          .     73:	numRequests.Add(1)
```


```bash
(pprof) peek encryptData
Showing nodes accounting for 65.08s, 100% of 65.08s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                             7.11s   100% |   main.encryptHandler
     0.01s 0.015% 0.015%      7.11s 10.93%                | main.encryptData
                                             6.79s 95.50% |   io.ReadFull (inline)
                                             0.13s  1.83% |   crypto/aes.NewCipher
                                             0.08s  1.13% |   runtime.makeslice
                                             0.04s  0.56% |   crypto/cipher.NewCFBEncrypter (inline)
                                             0.03s  0.42% |   crypto/cipher.(*cfb).XORKeyStream
                                             0.03s  0.42% |   encoding/hex.EncodeToString (inline)
----------------------------------------------------------+-------------
```

Other commands like `web` commands open the browser and `pdf` generates a PDF output

This may be the simplest of the code or something I am aware of but in this case we have intentionally uses an inefficient memory allocation strategy.

**LETS OPTIMIZE OUR CODE** 

Comment out the below lines
```go
// Intentionally inefficient: allocate new slice each iteration
encrypted := make([]byte, len(data))
stream.XORKeyStream(encrypted, data)
ciphertext = append(iv, encrypted...)  
```

And uncomment
```go
// More efficient: use the existing slice
// stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
```


### Re-Run profiling
Now if we run pprof again, we can see the improvements.
```
(pprof) list encryptData
Total: 57.98s
ROUTINE ======================== main.encryptData in /Users/vishal/go/src/github.com/vdparikh/goperf/go-perf/main.go
         0      6.51s (flat, cum) 11.23% of Total
         .          .     48:func encryptData(data []byte, key []byte) (string, error) {
         .      200ms     49:	block, err := aes.NewCipher(key)
         .          .     50:	if err != nil {
         .          .     51:		return "", err
         .          .     52:	}
         .       10ms     53:	ciphertext := make([]byte, aes.BlockSize+len(data))
         .          .     54:	iv := ciphertext[:aes.BlockSize]
         .      6.22s     55:	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
         .          .     56:		return "", err
         .          .     57:	}
         .          .     58:
         .       30ms     59:	stream := cipher.NewCFBEncrypter(block, iv)
         .          .     60:
         .          .     61:	// More efficient: use the existing slice
         .       10ms     62:	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
         .       40ms     63:	return hex.EncodeToString(ciphertext), nil
         .          .     64:
         .          .     65:}
         .          .     66:
         .          .     67:func encryptHandler(w http.ResponseWriter, r *http.Request) {
         .          .     68:	numRequests.Add(1)
```

### Let's Re-Run Benchmarking
Before
```
BenchmarkEncryptData-10               	 1195476	      1005 ns/op
```

After Optimizations
```
goos: darwin
goarch: arm64
pkg: github.com/vdparikh/goperf/profiling
BenchmarkEncryptData-10               	 1292698	       973.0 ns/op
BenchmarkEncryptDataSmall-10          	 1470314	       839.0 ns/op
BenchmarkEncryptDataMedium-10         	  367623	      3282 ns/op
BenchmarkEncryptDataLarge-10          	     508	   2274341 ns/op
BenchmarkEncryptDataAES256-10         	 1362688	       873.4 ns/op
BenchmarkEncryptDataConcurrency-10    	 1997667	       585.5 ns/op
PASS
ok  	github.com/vdparikh/goperf/profiling	11.062s
```

