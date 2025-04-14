package helpers

import (
	"fmt"
	"github.com/Weso1ek/go-gundam-hunter/context"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

func GetHtml(rawURL string) (string, error) {
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

	return string(htmlBodyBytes), nil
}

func ProcessHtml(htmlBody string, offers []*context.GundamOffers) []*context.GundamOffers {
	reader := strings.NewReader(htmlBody)

	parsed, err := html.Parse(reader)
	if err != nil {
		return offers
	}

	ProcessOffers(parsed, &offers)

	return offers
}

func ProcessOffers(n *html.Node, offers *[]*context.GundamOffers) {
	if n.Type == html.ElementNode && n.Data == "article" {
		fmt.Println("----")
		fmt.Println(n)
		fmt.Println("----")
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ProcessOffers(c, offers)
	}
}
