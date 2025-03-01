# authbench - HTTP Benchmarking Tool with basic authentication

This is a program for sending parallel HTTP requests with basic authentication to a specified URL. You can configure the authentication failure rate.

## Usage 

Run the program using the following command-line options:

```
Usage of ./authbench:
  -h string
        Target server (default "localhost")
  -uri string
        Request URI (default "/")
  -port int
        Port number (default 8080)
  -n int
        Number of requests to send (default 100)
  -c int
        Number of concurrent requests (default 10)
  -user string
        Username for basic authentication (default "admin")
  -pass string
        Password for basic authentication (default "password")
  -r float
        Authentication failure probability (0.0 - 1.0) (default 0.1)
  -i float
        Interval between repeated executions
  -R int
        Number of repetitions (default 1)
```

## Example 

The following command sends 100 requests in 10 parallel threads to a server running on localhost at port 8080, using basic authentication.

```
go run authbench.go -h localhost -p 8080 -user "testuser" -pass "testpass" -r 0.0 -R 100
```

## Output 

The program prints the status code and processing time for each request, along with the total execution time and the average execution time.

```
Starting HTTP Benchmarking...
Target URL: http://localhost:8080/
Requests: 100, Concurrency: 10, FailRate: 0.10, Repeat: 1, Interval: 0.0 sec

[Run 1/1] Starting benchmark...
[Worker 0] Status: 200 Time: 123.456µs
[Worker 1] Status: 200 Time: 234.567µs
...
Total time: 1.234567s

Benchmark completed. Average execution time: 1.234567s
```
