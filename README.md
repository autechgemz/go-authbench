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
$ go run authbench.go -h localhost -p 8080 -user "testuser" -pass "testpass" -r 0.0 -R 1 -i 1 -n 10 -c 10
Starting HTTP Benchmarking...
Target URL: http://localhost:8080/
Requests: 10, Concurrency: 10, FailRate: 0.00, Repeat: 1, Interval: 1.0 sec

[Run 1/1] Starting benchmark...
[Worker 7] Status: 200 Time: 5.248942ms
[Worker 0] Status: 200 Time: 5.707282ms
[Worker 9] Status: 200 Time: 6.129671ms
[Worker 3] Status: 200 Time: 6.127834ms
[Worker 4] Status: 200 Time: 6.117628ms
[Worker 6] Status: 200 Time: 6.211342ms
[Worker 2] Status: 200 Time: 7.422744ms
[Worker 8] Status: 200 Time: 7.506822ms
[Worker 5] Status: 200 Time: 7.547502ms
[Worker 1] Status: 200 Time: 7.301094ms
Total time: 7.679446ms

Benchmark completed. Average execution time: 7.679446ms
```
