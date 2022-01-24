package utils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func CheckHTTPResponse(urls string) {
	resp := ReadFile(urls)
	client := &http.Client{}

	for i := range resp {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		req, _ := http.NewRequest("GET", resp[i], nil)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[i] %d %s\n", resp.StatusCode, resp.Request.URL)
	}
}
