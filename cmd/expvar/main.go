package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

var (
	numRequests = expvar.NewInt("num_requests")
	stats       = expvar.NewMap("stats")
)

func init() {
	// Publish declares a named exported variable. This should be called from a package's init function when it creates its Vars. If the name is already registered then this will log.Panic.
	// Publish Metrics
	start := time.Now()
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

// PrintMemUsage print current memory consumption of Go
func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	log.Printf(
		"Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v",
		(m.Alloc / 1024 / 1024),
		(m.TotalAlloc / 1024 / 1024),
		m.Sys,
		m.NumGC,
	)
}

func handler(w http.ResponseWriter, r *http.Request) {
	numRequests.Add(1)
	stats.Add("handled", 1)

	printMemUsage()

	data := []byte("Hello, World! This is a test string for encryption.")
	w.Write([]byte(data))
}

func main() {
	// Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./cmd/expvar/index.html")
	})

	// http.HandleFunc("/debug/vars", expvar.Handler().ServeHTTP)

	http.HandleFunc("/handler", handler)

	http.ListenAndServe(":8080", nil)
}
