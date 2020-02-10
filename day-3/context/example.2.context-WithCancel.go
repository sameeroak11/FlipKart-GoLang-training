package main


import (
	"fmt"
	"context"
	"time"
	"os"
	"bufio"
	"runtime"
)


func myCron(ctx context.Context, timeAfter time.Duration, ctxCancelMsg string, cronMsg string) {
	select {
		case val := <-ctx.Done():
			fmt.Printf("ctx.Cancel: %s\n", ctxCancelMsg)
			fmt.Printf("ctx value: %#v\n", val)
			break
		case <-time.After(timeAfter):
			fmt.Println("cronMsg:", cronMsg)
			break
	}
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // allocates one logical processor for the scheduler to use

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		pS := bufio.NewScanner(os.Stdin)
		pS.Scan()
		// pReader := bufio.NewReader(os.Stdin)
		// pReader.ReadString('\n')
		cancel()
	}()

	myCron(ctx, 3 * time.Second, "context cancelled", "cronjob")
}
