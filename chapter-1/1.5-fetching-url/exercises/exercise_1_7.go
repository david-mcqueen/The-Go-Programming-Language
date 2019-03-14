package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Use io.Copy to copy the file, as opposed to ioutil.ReadAll

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		//b, err := ioutil.ReadAll(resp.Body)

		c, err := io.Copy(os.Stdout, resp.Body)


		resp.Body.Close()
		fmt.Printf("\f", c)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
