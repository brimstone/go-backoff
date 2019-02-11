package backoff

import (
	"context"
	"math/rand"
	"time"
)

// Inspired by https://upgear.io/blog/simple-golang-retry-function/
func init() {
	rand.Seed(time.Now().UnixNano())
}

func Backoff(ctx context.Context, b *time.Duration) {
	*b++
	// Add some jitter to the backoff time
	jitter := time.Duration(rand.Int63n(int64(*b)))
	*b = *b + jitter/2
	// Create a ticker to wait
	ticker := time.NewTicker(*b)
	// Make sure to free those resources
	defer ticker.Stop()

	select {
	// If the context is done before the ticker, return
	case <-ctx.Done():
		return
		// if the ticker is done before the context, return
	case <-ticker.C:
		return
	}
}
