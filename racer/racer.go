package racer

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const DEFAULT_TIMEOUT = 10 * time.Second

// Ping sends an HTTP GET request to the specified URL and returns the response status.
// If the request is successful (status code 200), it returns "OK".
// If there is an error during the request, it returns the error message.
func Ping(url string) string {
	log.Printf("Pinging %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}

	log.Printf("Response Headers: %v", resp.Header)
	log.Printf("Response Status: %v", resp.Status)
	log.Printf("Response Body: %v", resp.Body)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return "OK"
	}

	return resp.Status
}


// TimedPing measures the time it takes to ping a specified URL.
// It returns the duration of the ping in nanoseconds.
func TimedPing(url string) time.Duration {
	start := time.Now()

	Ping(url)

	return time.Since(start)	
}

// Racer compares the response times of two URLs and returns the URL with the faster response time.
// It uses the TimedPing function to measure the response times.
func Racer(a, b string) string {
	aDuration := TimedPing(a)
	bDuration := TimedPing(b)

	if aDuration < bDuration {
		return a
	}

	return b
}


// PingAsync sends an asynchronous ping request to the specified URL.
// It returns a channel that will be closed when the ping request is complete.
func PingAsync(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		Ping(url)
		close(ch)
	}()

	return ch
}


// RacerAsyncWithTimeout compares the response times of two URLs asynchronously with a timeout.
// It returns the URL that responds first, or an error if both URLs timeout.
func RacerAsyncWithTimeout(a, b string,  timeout time.Duration) (string, error) {
	select {
	case <-PingAsync(a):
		return a, nil
	case <-PingAsync(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// RacerAsync compares the response times of two URLs asynchronously and returns the URL that responds first.
// It uses a default timeout value if no timeout is specified.
func RacerAsync(a, b string) (string, error) {
	return RacerAsyncWithTimeout(a, b, DEFAULT_TIMEOUT)
}