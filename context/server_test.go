package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

// Fetch is a method of the SpyStore struct that fetches data.
// It takes a context.Context as an argument and returns a string and an error.
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {

	// Create a buffered channel to store the result string with a capacity of 1.
	data := make(chan string, 1)

	// Start a goroutine to fetch the data asynchronously.
	go func() {
		// Initialize an empty string to store the result.
		var result string
		// Iterate over the characters in the response slice of the SpyStore.
		for _, c := range s.response {
			// Use a select statement to perform non-blocking operations.
			select {
			// If the context is canceled, print a message and return.
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			// If the context is not canceled, continue fetching data.
			default:
				// Simulate some processing time.
				time.Sleep(10 * time.Millisecond)
				// Append the character to the result string.
				result += string(c)
			}
		}
		// Send the result string to the data channel.
		data <- result
	}()

	// Use a select statement to wait for either the context to be canceled or data to be available.
	select {
	// If the context is canceled, return an empty string and the context error.
	case <-ctx.Done():
		return "", ctx.Err()
	// If data is available, receive it from the data channel and return it along with nil error.
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("Returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	// This test function verifies whether the server cancels its work
	// if the HTTP request is cancelled.
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		// Define some test data.
		data := "Hello World"
		// Create a SpyStore instance with the test data and the testing context.
		store := &SpyStore{response: data, t: t}
		// Initialize the server with the SpyStore.
		svr := Server(store)

		// Create a new HTTP request for a GET method with a "/" endpoint.
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// Create a new context that can be cancelled and a function to cancel it.
		cancellingCtx, cancel := context.WithCancel(request.Context())
		// Schedule the cancellation of the context after 5 milliseconds.
		time.AfterFunc(5*time.Millisecond, cancel)
		// Associate the cancelling context with the request.
		request = request.WithContext(cancellingCtx)

		// Create a SpyResponseWriter instance to capture the response.
		response := &SpyResponseWriter{}

		// Serve the HTTP request using the server.
		svr.ServeHTTP(response, request)

		// Check if the response has been written.
		if response.written {
			// If the response has been written, it indicates an error.
			t.Error("A response should not have been written")
		}
	})
}
