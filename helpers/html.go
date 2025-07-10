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

func ProcessHtml(htmlBody string, offers []*context.GundamOffer) []*context.GundamOffer {
	reader := strings.NewReader(htmlBody)

	parsed, err := html.Parse(reader)
	if err != nil {
		return offers
	}

	ProcessOffers(parsed, &offers)

	return offers
}

func ProcessOffers(n *html.Node, offers *[]*context.GundamOffer) {
	if n.Type == html.ElementNode && n.Data == "article" {
		processOffer(n, offers)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ProcessOffers(c, offers)
	}
}

func processOffer(n *html.Node, offers *[]*context.GundamOffer) {

	gundamOffer := &context.GundamOffer{}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "div" {
			for d := c.FirstChild; d != nil; d = d.NextSibling {
				if d.Data == "a" {
					for _, a := range d.Attr {
						if a.Key == "href" {
							gundamOffer.Url = a.Val
						}
					}

					for f := d.FirstChild; f != nil; f = f.NextSibling {
						for _, fa := range f.Attr {
							if fa.Key == "src" {
								gundamOffer.Img = fa.Val
							}
						}
					}
				}

				if d.Data == "div" {
					for e := d.FirstChild; e != nil; e = e.NextSibling {
						if e.Data == "h2" {

						}
					}
				}
			}
		}
	}

	*offers = append(*offers, gundamOffer)
}
