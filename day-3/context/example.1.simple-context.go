package main


import (
	"fmt"
	"context"
	"time"
)


func myCron(ctx context.Context, schTime time.Duration, ctxCancelMsg string, cronMsg string) {
	select {
		case <-ctx.Done():
			fmt.Println("ctx.Cancel:", ctxCancelMsg)
			break
		case <-time.After(schTime):
			fmt.Println("cronMsg:", cronMsg)
			break
	}
}


func main() {
	ctx := context.Background()
	sleepAndTalk(ctx, 5 * time.Second, "hello.1")
}
