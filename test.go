package main
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func test()() {
	for {
		fmt.Println("hi")
	}
}
func main() {
    cancelChan := make(chan os.Signal, 1)
    // catch SIGETRM or SIGINTERRUPT
    signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
    go func() {
        // start your software here. Maybe your need to replace the for loop with other code


            // replace the time.Sleep with your code
            test()

    }()
    sig := <-cancelChan
    fmt.Printf("Caught signal %v", sig)
    // shutdown other goroutines gracefully
    // close other resources
}