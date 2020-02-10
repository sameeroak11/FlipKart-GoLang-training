package main


import (
	"fmt"
	"context"
	"time"
	"runtime"
)


func myCron(ctx context.Context, ctxCancelMsg string, timeout time.Duration, cronMsg string) {
	select {
		case <-ctx.Done():
			fmt.Println("ctx.Cancel WithDeadline:", ctxCancelMsg)
			break
		/* case <-time.After(timeout):
			fmt.Println("cronMsg:", cronMsg)
			break */
	}
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // allocates one logical processor for the scheduler to use

	ctx := context.Background()
	//timeStamp, _ := time.Parse("02-01-2006 15:04:05", time.Now().Local().Format("02-01-2006 15:04:05"))
	//ctx, contextCancel := context.WithDeadline(ctx, time.Now().Local().Add(3 * time.Second))

	timeStamp, _ := time.Parse(time.RFC3339, "2020-01-18T18:34:40+05:30")
	//ctx, contextCancel := context.WithDeadline(ctx, timeStamp.Add(5 * time.Second))
	ctx, contextCancel := context.WithDeadline(ctx, timeStamp)
	defer contextCancel()

	myCron(ctx, "context cancelled", 10 * time.Second, "cronjob")
}
