package racer

import (
	"net/http/httptest"
	"testing"
	"time"
	"net/http"
)

func TestRacer(t *testing.T) {
	t.Run("Ping gets a url", func(t *testing.T) {
		url := "http://www.google.com"

		got := Ping(url)
		want := "OK"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Ping gets a url that doesn't exist", func(t *testing.T) {
		url := "http://www.google.com/404"

		got := Ping(url)
		want := "404 Not Found"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("TimedPing", func(t *testing.T) {
		url := "http://www.google.com"

		got := TimedPing(url)
		if got >= time.Second {
			t.Errorf("got %v, want less than 1 second", got)
		}
	})


	t.Run("Racer", func(t *testing.T) {
		slowServer := createMockServer(20 * time.Millisecond)
		defer slowServer.Close()
		fastServer := createMockServer(0)
		defer fastServer.Close()

		got := Racer(slowServer.URL, fastServer.URL)
		want := fastServer.URL
		
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("RacerAsync", func(t *testing.T) {
		slowServer := createMockServer(20 * time.Millisecond)
		defer slowServer.Close()
		fastServer := createMockServer(0)
		defer fastServer.Close()

		got, err := RacerAsyncWithTimeout(slowServer.URL, fastServer.URL, 30 * time.Millisecond)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		want := fastServer.URL
		
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("RacerAsync return an error if no servers responds within Timeout", func(t *testing.T) {
		serverA := createMockServer(11 * time.Millisecond)
		serverB := createMockServer(12 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := RacerAsyncWithTimeout(serverA.URL, serverB.URL, 1 * time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func createMockServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}