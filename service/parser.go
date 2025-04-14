package service

import (
	"fmt"
	"github.com/Weso1ek/go-gundam-hunter/context"
	"github.com/Weso1ek/go-gundam-hunter/helpers"
)

func GetOffers(searchTerm string) []*context.GundamOffers {
	var offers []*context.GundamOffers

	for _, k := range context.SearchServices {
		fmt.Println("Parsing: ", k.Domain)

		serviceHtml, err := helpers.GetHtml(k.SearchUrl + searchTerm)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(serviceHtml)
	}

	return offers
}
