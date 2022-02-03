package main

import (
        "fmt"
		"io"
        "os"
		// "time" // https://pkg.go.dev/time
		"encoding/json"
)

func main() {
        fmt.Println("Print from the Go program")
        fmt.Println(os.Getenv("TEST_ENV"))

		io.WriteString(os.Stdout,"This is the line to standard output.\n")
		io.WriteString(os.Stderr,"This is the line for standard error output.\n")

		// --- time output constant --- //
		// print every 5 seconds how long the program is running
		// for range time.Tick(time.Second * 30) {
		// 	go func() {
		// 		fmt.Println(os.Stdout, time.Now())
		// 	}()
		// }

		// --- JSON --- //
		data_log := map[string]interface{}{
			// "@timestamp": 2022-02-01T08:50:02.711889+00:00,
			"@source": "b58a81938fbf",
			"@fields": "",
			"@message": "Api request",
			// "@tags": ["api-request-logger"],
			"@type": "api-request-response",
		}


		// https://go.dev/play/p/LY0RWyrT6r4
		res := map[string]interface{}{
			"one": 1,
			"two": "twotwo",
		}
		if err := json.NewEncoder(os.Stdout).Encode(res); err != nil {
			panic(err)
		}
		// output:
		// {"one":1,"two":"twotwo"}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(res); err != nil {
			panic(err)
		}
		// output:
		// {
		// 	"one": 1,
		// 	"two": "twotwo"
		// }
}

