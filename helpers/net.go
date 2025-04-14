package helpers

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("got network error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("got HTTP error: %s", res.Status)
	}

	htmlBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}

	htmlBody := string(htmlBodyBytes)

	return htmlBody, nil
}
