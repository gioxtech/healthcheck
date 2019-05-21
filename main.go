package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	httpAddrFlag = flag.String("http-addr", "", "http addr to check if service is health")
)

func init() {
	flag.Parse()
}

func main() {
	if httpAddr := *httpAddrFlag; httpAddr != "" {
		resp, err := http.Get(httpAddr)
		if err != nil {
			fmt.Println("Unable to perform health check:", err.Error())
			os.Exit(1)
		}

		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)

		if resp.StatusCode > 299 {
			fmt.Println("Service is not healthy, received:", http.StatusText(resp.StatusCode))
			os.Exit(1)
		}
		os.Exit(0)
	}

	flag.Usage()
	os.Exit(1)
}
