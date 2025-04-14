package service

import (
	"fmt"
	"github.com/Weso1ek/go-gundam-hunter/context"
	"github.com/Weso1ek/go-gundam-hunter/helpers"
	"net/url"
)

func GetOffers(searchTerm string) []*context.GundamOffers {
	var offers []*context.GundamOffers

	for _, k := range context.SearchServices {
		searchUrl := k.SearchUrl + url.QueryEscape(searchTerm)

		fmt.Println("Parsing: ", k.Domain)
		fmt.Println("Parsing: ", searchUrl)

		serviceHtml, err := helpers.GetHtml(searchUrl)
		if err != nil {
			fmt.Println(err)
		}

		domainOffers := helpers.ProcessHtml(serviceHtml, offers)

		fmt.Println(domainOffers)
	}

	return offers
}
