package main

import (
	"context"
	"fmt"
	"time"
)

// START MAIN OMIT
func work(ctx context.Context, reason string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Finished", reason, "work", ctx.Value("foo"))
			return
		default:
			fmt.Println("Doing", reason, "work...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx, "cancel")
	time.Sleep(3 * time.Second)
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	ctx = context.WithValue(ctx, "foo", "bar")
	go work(ctx, "timeout")
	time.Sleep(4 * time.Second)
}

// END MAIN OMIT
