# Getting Started with Benchmarking and Go
This doc will walk you through basic benchmarking in Go

The full code can be found in `cmd/benchmarking/main.go`


Now that you know how to do monitoring using `expvar` and load test using `hey`, lets look at Benchmarking.

Benchmarking in Golang is a systematic way to measure the performance of your code using the testing package. Benchmarks can provide valuable insights into how different parts of your code perform under certain conditions and can help you identify potential areas for optimization.

### Running the Benchmark
Benchmarks are written similar to tests, but use b *testing.B to provide a loop count b.N, which adjusts until the framework is confident in the stability of the results.

```go
import (
	"testing"
)

func BenchmarkEncryptData(b *testing.B) {
	data := []byte("Hello, World! This is a test string for encryption.")
	key := generateKey()
	for i := 0; i < b.N; i++ {
		_, err := encryptData(data, key)
		if err != nil {
			b.Fatal(err)
		}
	}
}
```
Take a look at the `cmd/benchmarking/main_test.go` file to see more benchmark tests. 

To execute Benchmark tests, you can use `cd cmd/benchmarking/ && go test -bench=.`

Once the test completes, you will see an output like 
```bash
goos: darwin
goarch: arm64
pkg: github.com/vdparikh/goperf/profiling
BenchmarkEncryptData-10               	 1195476	      1005 ns/op
BenchmarkEncryptDataSmall-10          	 1384868	       859.1 ns/op
BenchmarkEncryptDataMedium-10         	  336432	      3533 ns/op
BenchmarkEncryptDataLarge-10          	     470	   2613281 ns/op
BenchmarkEncryptDataAES256-10         	 1302998	       944.2 ns/op
BenchmarkEncryptDataConcurrency-10    	 1933846	       653.8 ns/op
PASS
ok  	github.com/vdparikh/goperf/profiling	11.334s
```

You can add additional parameters like `-benchmem` and `benchtime`. Benchmem allows you to see memory allocation like `go test -run none -bench . -benchtime 3s -benchmem`

which provides extended output like
```
BenchmarkEncryptData-10               	 3835315	       928.0 ns/op	     992 B/op	      11 allocs/op
BenchmarkEncryptDataSmall-10          	 4424530	       804.4 ns/op	     736 B/op	      11 allocs/op
```

The 992 B/op shows 992 bytes allocated over one object. Final column shows 11 objects on the heap worth 992 bytes of memory.

### Interpreting Results
The values in the output
- 1195476 is the number of times the function was executed.
- 1005 ns/op means each operation took on average 1005 nanoseconds.

_Results include the number of operations and the time per operation. It's critical to run benchmarks in a clean environment to avoid fluctuations caused by other processes_

- **Time per Operation**: The ns/op metric tells you how long, on average, each operation took to complete. A lower number indicates better performance.
- **Comparative Benchmarking**: Running benchmarks before and after changes in your code can show you how those changes affect performance. We will take a look at this soon!
- **Resource Usage**: Along with the execution time, consider other resources like memory and CPU usage. These can be profiled separately to understand the impact of your code changes comprehensively.

### Profiling with Benchmark Data
```
go test -bench=. -cpuprofile cpu.out -memprofile mem.out
go tool pprof cpu.out
go tool pprof mem.out
```

What is Profiling? Glad you asked and so lets take a look at the next on getting-started-with-profiling.
