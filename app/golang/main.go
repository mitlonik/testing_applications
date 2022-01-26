package main

import (
        "fmt"
		"io"
        "os"
		"time" // https://pkg.go.dev/time
)

func main() {
        fmt.Println("Print from the Go program")
        fmt.Println(os.Getenv("TEST_ENV"))

		io.WriteString(os.Stdout,"This is the line to standard output.\n")
		io.WriteString(os.Stderr,"This is the line for standard error output.\n")

		// print every 5 seconds how long the program is running
		for range time.Tick(time.Second * 5) {
			go func() {
				fmt.Println(os.Stdout, time.Now())
			}()
		}
}
