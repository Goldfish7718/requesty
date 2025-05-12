package utils

import (
	"io"
	"log"
	"net/http"
)

func ParseResponseBody(resp *http.Response) string {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	return string(body)
}
