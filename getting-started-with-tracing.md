# Getting Started with Tracing and Go
This doc will walk you through basic tracing in Go

The full code can be found in `cmd/tracing/main.go`

https://blog.gopheracademy.com/advent-2017/go-execution-tracer/

Go's trace tool to analyze performance issues in your encryption example can provide a deep dive into how concurrency and system calls behave under load.  You can use Go package `runtime/trace` for tracing. 

- The tracer is a powerful tool for debugging concurrency issues, e.g, contentions and logical races. But it does not solve all problems: it is not the best tool available to track down what piece of code is spending most CPU time or allocations. The go tool pprof is better suited for these use cases.
- Go's trace tool collects event data about your Go program's execution and visualizes it with the help of the Chrome tracing tool. It's particularly useful for understanding system scheduling, synchronization details, and network-related delays in Go programs.

Set up tracing in your main function. You'll start by creating a file where the trace output will be stored. Then, enable tracing with trace.Start and ensure it stops with defer.

```go
import(
  "runtime/trace"
  // ...
)

func main() {
	// Enable Tracing
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

  // .... Rest of the code
  // ...

  // http.ListenAndServe blocks until it's not listening to the port (for example if there was an error binding to the port). If it's not run in a separate goroutine, subsequent code (in this case trace.Stop()) will never get called
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Exit Signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}

```
Once you run the application and put some load (`hey -n 1000 -c 100 -z 30s http://localhost:8080/encrypt`) so the tracer can collect meaningful data.
The output of the trace is written to trace.out.

You can now run the trace.out with trace tool to see the results

```
go tool trace trace.out
```

This command will open up a web interface in your browser where you can visually analyze the trace. The tool provides various views that show the timeline of events, goroutine blocking profiles, network/synchronization blocking, and more.

## Interpreting Results
- View Goroutine Activities: Helps in identifying delays and potential deadlocks.
- Network/Syscall Blocking: Shows how often and why goroutines block on network or syscall operations, useful for identifying I/O bottlenecks.
- Synchronization Blocking: Useful for spotting contention issues around mutexes or other synchronization primitives.
