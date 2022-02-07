package main

import (
        "fmt"
		"io"
        "os"
		"time" 			// https://pkg.go.dev/time
)

func main() {
		// --- time output constant --- //
		// print every 5 seconds how long the program is running
		for range time.Tick(time.Second * 30) {
			go func() {
				fmt.Println(os.Stdout, time.Now())
			}()
		}
}

