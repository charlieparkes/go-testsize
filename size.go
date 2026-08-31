// go-testsize implements declarative methods for go tests which enforce
// test runtime limits using context timeouts. Test sizes and timeouts
// adhere to the Google Test Taxonomy.
//
//	+-----------------------+-------+----------------+-------+
//	| Feature               | Small | Medium         | Large |
//	+-----------------------+-------+----------------+-------+
//	| Network access        | No    | localhost only | Yes   |
//	| Database              | No    | Yes            | Yes   |
//	| File system access    | No    | Yes            | Yes   |
//	| Use external systems  | No    | Discouraged    | Yes   |
//	| Multiple threads      | No    | Yes            | Yes   |
//	| Sleep statements      | No    | Yes            | Yes   |
//	| System properties     | No    | Yes            | Yes   |
//	| Time limit (seconds)  | 60    | 300            | 900+  |
//	+-----------------------+-------+----------------+-------+
//
// https://testing.googleblog.com/2010/12/test-sizes.html
package testsize

import (
	"context"
	"testing"
	"time"
)

// Timeouts follow Google's test size taxonomy.
const (
	SmallTimeout  = time.Minute
	MediumTimeout = 5 * time.Minute
	LargeTimeout  = 15 * time.Minute
)

// Small returns a context that times out after SmallTimeout.
func Small(tb testing.TB) context.Context {
	tb.Helper()
	return withTimeout(tb, SmallTimeout)
}

// Medium returns a context that times out after MediumTimeout.
func Medium(tb testing.TB) context.Context {
	tb.Helper()
	return withTimeout(tb, MediumTimeout)
}

// Large returns a context that times out after LargeTimeout.
func Large(tb testing.TB) context.Context {
	tb.Helper()
	return withTimeout(tb, LargeTimeout)
}

func withTimeout(tb testing.TB, timeout time.Duration) context.Context {
	tb.Helper()
	ctx, cancel := context.WithTimeout(tb.Context(), timeout)
	tb.Cleanup(cancel)
	return ctx
}
