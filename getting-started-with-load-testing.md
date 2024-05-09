# Getting Started with Load Testing and Go
This doc will walk you through basic benchmarking in Go

The full code can be found in `cmd/loadtesting/main.go`

Load testing helps simulates real-world load on your application and helps understand how they behave under stress. This helps to identify the maximum operating capacity of the app, including its scalability and points of failure.

There are several open source perfomance tools like arc, JMeter, Locust, k6, Apache Bench. Well you can even just do a normal shell script load on a endpoint by doing a for loop.
```bash
for i in {1..100}; do
  curl http://localhost:8080/encrypt &
done
wait
```

But for more robust testing and results, we can use a more sophistacted package like `hey` (github.com/rakyll/hey).

`hey` is a modern load testing tool written in Go, designed to send a high volume of requests to a target web application in order to simulate traffic and measure performance. It's a lightweight and easy-to-use tool that's particularly useful for testing HTTP services. 

You can install it on your mac using brew `brew install hey`. 

Once you install hey, you can run your load test against the API endpoint using `hey -options url`

Here are some of the most commonly used options:
- -n: Number of requests to run. Default is 200.
- -c: Number of workers to run concurrently. Each worker will make -n requests. Default is 50.
- -q: Rate limit, in queries per second (QPS) per worker. Default is no rate limit.
- -z: Duration of the test, e.g., 10s for 10 seconds. This can be used instead of -n for a duration-based test.
- -m: HTTP method, default is GET.
- -h: Custom HTTP header. Example: -h "Authorization: Bearer <token>". You can specify as many headers as you need by repeating the -h flag.
- -d: HTTP request body data.
- -T: Content type of the body data.

```bash
# Call /encrypt endpoint with 1000 total requests, 100 concurrent workers, over a period of 30 seconds
hey -n 1000 -c 100 -z 30s http://localhost:8080/encrypt
```

### Interpreting Results
```
Summary:
  Total:	5.0018 secs
  Slowest:	0.0289 secs
  Fastest:	0.0000 secs
  Average:	0.0014 secs
  Requests/sec:	71554.8622

  Total data:	47959270 bytes
  Size/request:	134 bytes

Response time histogram:
  0.000 [1]	|
  0.003 [324380]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [28602]	|■■■■
  0.009 [3752]	|
  0.012 [734]	|
  0.014 [223]	|
  0.017 [97]	|
  0.020 [96]	|
  0.023 [6]	|
  0.026 [3]	|
  0.029 [11]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0006 secs
  50% in 0.0011 secs
  75% in 0.0016 secs
  90% in 0.0028 secs
  95% in 0.0037 secs
  99% in 0.0064 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0289 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0041 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0131 secs
  resp wait:	0.0013 secs, 0.0000 secs, 0.0270 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0159 secs

Status code distribution:
  [200]	357905 responses
```

hey provides a summary of the test results once it completes, which includes:

- Total requests: Total number of requests made.
- Total duration: Total time taken for the test.
- Response time statistics: Includes average, fastest, slowest, and times across percentiles (e.g., 50th, 90th, 95th).
- Requests per second: The average number of requests per second.
- Total data transferred: Amount of data transferred during the test.
- Status code distribution: Counts of the returned HTTP status codes.

### Advanced Usage
You can also use hey to perform more complex load testing scenarios such as:

- Testing different HTTP methods: Use -m POST with -d @data.json to send POST requests with data from a file.
- Headers: Use multiple -h flags to include custom headers, such as authentication tokens.
- Rate limiting: If you want to simulate a more controlled load, use the -q option to specify a rate limit.

### Automation and Integration
hey can be integrated into scripts for regular performance testing, or as part of CI/CD pipelines to ensure performance benchmarks are met before deploying to production.

