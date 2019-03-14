package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Append http:// as a prefix, if it doesnt already exist

func main() {
	for _, url := range os.Args[1:] {

		urlToUse := url

		if !strings.HasPrefix(urlToUse, "http://") {
			fmt.Fprintf(os.Stderr, "fetch: Adding Prefix\n")
			urlToUse = "http://" + urlToUse
		}

		resp, err := http.Get(urlToUse)
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
