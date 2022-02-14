package main

import (
        // "fmt"
		"io"
        "os"
		"time" 			// https://pkg.go.dev/time
		"encoding/json" // https://go.dev/play/p/LY0RWyrT6r4
)

func main() {

		io.WriteString(os.Stdout,"This is the line to standard output.\n")
		io.WriteString(os.Stderr,"This is the line for standard error output.\n")


		// for range time.Tick(time.Second * 10) {
		// 	go func() {
		// 		fmt.Println(os.Stdout, time.Now())
		// 	}()
		// }

		// --- JSON --- //
		data_log := map[string]interface{}{
			"@source": "b58a81938fbf",
			"@fields": "",
			"@message": "Api request",
			"@type": "api-request-response",
		}

		for range time.Tick(time.Second * 5) {
			go func() {
				if err := json.NewEncoder(os.Stdout).Encode(data_log); err != nil {
					panic(err)
				}
			}()
		}
}