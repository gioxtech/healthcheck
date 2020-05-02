package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	httpAddrFlag = flag.String("http-addr", "", "http addr to check if service is health")
	insecureFlag = flag.Bool("insecure", false, "if invalid certificate should be ignored or not")
)

func init() {
	flag.Parse()
}

func main() {
	if httpAddr := *httpAddrFlag; httpAddr != "" {
		if *insecureFlag {
			httpTransport := http.DefaultTransport.(*http.Transport)
			httpTransport.TLSClientConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}

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
