package main

import (
	"context"
	"fmt"
)

// always use unexported types for keys to avoid collisions
type ctxKey string

func a(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey("userID"), 42)
}

func b(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey("requestID"), "req-123")
}

func main() {
	ctx := context.Background()

	ctx = a(ctx)
	ctx = b(ctx)

	fmt.Println(ctx.Value(ctxKey("userID")))    // 42
	fmt.Println(ctx.Value(ctxKey("requestID"))) // req-123
}
