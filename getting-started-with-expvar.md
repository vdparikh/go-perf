# Getting Started with Debugging and Go
This doc will walk you through basic debugging and monitoring in Go using `expvar`

The full code can be found in `cmd/expvar/main.go`

`expvar` is a package in Go that provides a standardized way to export live debugging and monitoring data from your applications in the form of variables. While it might not be as widely discussed as other performance and monitoring tools, it's still very useful, especially for long-running applications that need constant monitoring or debugging support. It's particularly handy because it integrates directly with Go's net/http package to serve up live application data in a web-accessible format, typically JSON.

### Why Use expvar?
- Lightweight Monitoring: Provides a simple and lightweight way to monitor internal variables.
- Ease of Integration: Easy to integrate with existing Go web applications.
- Real-Time Metrics: Allows access to real-time metrics via HTTP.

### How to use it?
It as simple as adding `"expvar"` import to your code. The `init()` function in the package takes care of exposing the `http.HandleFunc("/debug/vars", expvarHandler)` handler that will print a whole lot of information like GC, memory, commandline stats and custom metrics as a JSON document.

Creating custom metrics is also quite simple

```go
var (
	numRequests = expvar.NewInt("num_requests")
	stats       = expvar.NewMap("stats")
)

func init() {
  // Publish declares a named exported variable. This should be called from a package's init function when it creates its Vars. If the name is already registered then this will log.Panic.
  // Publish Metrics
  start := time.Now()

  // "uptime" metric. How long has my service been up? Using the traditional static approach, you would need to update this value periodically.
	expvar.Publish("uptime", expvar.Func(func() interface{} {
		return time.Since(start).Seconds()
	}))
  expvar.Publish("goroutines", expvar.Func(func() interface{} {
		return fmt.Sprintf("%d", runtime.NumGoroutine())
	}))
	expvar.Publish("cgocall", expvar.Func(func() interface{} {
		return fmt.Sprintf("%d", runtime.NumCgoCall())
	}))
	expvar.Publish("cpu", expvar.Func(func() interface{} {
		return fmt.Sprintf("%d", runtime.NumCPU())
	}))
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
  // Increment number of requests
	numRequests.Add(1)
  // Add to Map
	stats.Add("handled", 1)

  // ... Rest of the logic
}
```

### Viewing Metrics
You can now access `http://localhost:8080/debug/vars` to see all metrics. Its a big JSON document.
However if you want to build you own, head on to a simple example at `http://localhost:8080`

Alternatively you can use `expvarmon`. You can read more about the package at `https://github.com/divan/expvarmon`

```bash
go install github.com/divan/expvarmon@latest
expvarmon -ports="8080" #-vars="goroutines,num_requests,uptime"
```

### Best Practices
**Security Considerations:** It is important to secure the /debug/vars endpoint, especially in production environments. We can leverage below techniques to prevent access.

- Basic authentication
- Restricting IP access
- Only enabling expvar in development or via internal networks

**Performance Implications:** Note that while expvar is lightweight, indiscriminate use of logging or monitoring can impact application performance.

### Note 
- expvar primarily exposes key/value pairs
- expvar does not directly support arrays or slices because its main use is for exposing simple, thread-safe metrics.


