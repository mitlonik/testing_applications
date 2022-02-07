package main

import (
        "fmt"
		"io"
        "os"
		"encoding/json" // https://go.dev/play/p/LY0RWyrT6r4
)

func main() {

		// --- set vatiables --- //
		data_log := map[string]interface{}{
			"@source": "b58a81938fbf",
			"@fields": "",
			"@message": "Api request",
			"@type": "api-request-response",
		}

		// --- json output --- //
		if err := json.NewEncoder(os.Stdout).Encode(data_log); err != nil {
			panic(err)
		}
		// output:
		// {"one":1,"two":"twotwo"}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(data_log); err != nil {
			panic(err)
		}
		// output:
		// {
		// 	"one": 1,
		// 	"two": "twotwo"
		// }
}

