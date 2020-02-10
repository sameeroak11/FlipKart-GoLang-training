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
			fmt.Println("ctx.Cancel WithTimeout:", ctxCancelMsg)
			break
		case <-time.After(timer):
			fmt.Println("cronMsg:", cronMsg)
			break
	}
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // allocates one logical processor for the scheduler to use

	ctx := context.Background()
	ctx, contextCancel := context.WithTimeout(ctx, 1 * time.Second)
	defer contextCancel()

	myCron(ctx, "context cancelled", 3 * time.Second, "cronjob")
}
