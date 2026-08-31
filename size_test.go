package testsize_test

import (
	"context"
	"testing"
	"time"

	"github.com/charlieparkes/go-testsize"
)

func TestSizeDeadlines(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	tests := []struct {
		name    string
		size    func(testing.TB) context.Context
		timeout time.Duration
	}{
		{name: "small", size: testsize.Small, timeout: testsize.SmallTimeout},
		{name: "medium", size: testsize.Medium, timeout: testsize.MediumTimeout},
		{name: "large", size: testsize.Large, timeout: testsize.LargeTimeout},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			before := time.Now()
			ctx := tt.size(t)
			after := time.Now()

			deadline, ok := ctx.Deadline()
			if !ok {
				t.Fatal("context has no deadline")
			}

			minDeadline := before.Add(tt.timeout)
			maxDeadline := after.Add(tt.timeout)
			if deadline.Before(minDeadline) || deadline.After(maxDeadline) {
				t.Fatalf("deadline = %v, want between %v and %v", deadline, minDeadline, maxDeadline)
			}
		})
	}
}
