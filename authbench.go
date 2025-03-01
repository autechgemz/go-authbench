package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func generateBasicAuth(username, password string, failRate float64) string {

	if randomFail(failRate) {
		username = "wronguser"
		password = "wrongpass"
	}
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

}

func randomFail(rate float64) bool {
	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		return false
	}
	return float64(b[0])/255.0 < rate
}

func benchmark(url string, requests, concurrency int, username, password string, failRate float64) time.Duration {
	var wg sync.WaitGroup
	reqChan := make(chan int, requests)
	results := make(chan string, requests)
	startTime := time.Now()

	for i := 0; i < requests; i++ {
		reqChan <- i
	}
	close(reqChan)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			client := &http.Client{}
			for range reqChan {
				reqStart := time.Now()
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					results <- fmt.Sprintf("[Worker %d] Error: %v", workerID, err)
					continue
				}
				req.Header.Set("Authorization", generateBasicAuth(username, password, failRate))
				resp, err := client.Do(req)
				duration := time.Since(reqStart)
				if err != nil {
					results <- fmt.Sprintf("[Worker %d] Error: %v", workerID, err)
					continue
				}
				results <- fmt.Sprintf("[Worker %d] Status: %d Time: %v", workerID, resp.StatusCode, duration)
				resp.Body.Close()
			}
		}(i)
	}

	go func() {
		for res := range results {
			fmt.Println(res)
		}
	}()
	wg.Wait()
	close(results)

	totalTime := time.Since(startTime)
	fmt.Printf("Total time: %v\n", totalTime)
	return totalTime

}

func main() {
	server := flag.String("h", "localhost", "Server host")
	uri := flag.String("uri", "/", "Request URI")
	port := flag.Int("p", 8080, "Server port")
	requests := flag.Int("n", 100, "number of requests")
	concurrency := flag.Int("c", 10, "number of concurrent requests")
	username := flag.String("user", "admin", "basic auth username")
	password := flag.String("pass", "password", "basic auth password")
	failRate := flag.Float64("r", 0.1, "Failure rate")
	interval := flag.Float64("i", 0, "interval")
	repeat := flag.Int("R", 1, "bench count")
	flag.Parse()

	finalURL := "http://" + *server + ":" + strconv.Itoa(*port) + *uri

	fmt.Println("Starting HTTP Benchmarking...")
	fmt.Printf("Target URL: %s\n", finalURL)
	fmt.Printf("Requests: %d, Concurrency: %d, FailRate: %.2f, Repeat: %d, Interval: %.1f sec\n",
		*requests, *concurrency, *failRate, *repeat, *interval)

	var totalDurations time.Duration
	for i := 1; i <= *repeat; i++ {
		fmt.Printf("\n[Run %d/%d] Starting benchmark...\n", i, *repeat)
		runDuration := benchmark(finalURL, *requests, *concurrency, *username, *password, *failRate)
		totalDurations += runDuration
		if i < *repeat {
			sleepDuration := time.Duration(*interval*1000) * time.Millisecond
			fmt.Printf("[Run %d/%d] Sleeping for %.1f seconds...\n", i, *repeat, *interval)
			time.Sleep(sleepDuration)
		}
	}
	avgDuration := totalDurations / time.Duration(*repeat)
	fmt.Printf("\nBenchmark completed. Average execution time: %v\n", avgDuration)
}
