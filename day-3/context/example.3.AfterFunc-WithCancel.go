package main


import (
	"fmt"
	"context"
	"time"
	"runtime"
)


func myCron(ctx context.Context, ctxCancelMsg string, timer time.Duration, cronMsg string) {
	select {
		case <-ctx.Done():
			fmt.Println("ctx.Cancel WithCancel:", ctxCancelMsg)
			break
		case <-time.After(timer):
			fmt.Println("cronMsg:", cronMsg)
			break
	}
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // allocates one logical processor for the scheduler to use

	ctx := context.Background()
	ctx, contextCancel := context.WithCancel(ctx)

	time.AfterFunc(5 * time.Second, contextCancel)

	myCron(ctx, "context cancelled", 3 * time.Second, "cronjob")
}
