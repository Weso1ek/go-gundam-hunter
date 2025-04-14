package context

type GundamOffers struct {
	Title  string
	Price  float64
	Link   string
	Domain string
	Img    *string
}

type SearchService struct {
	Domain    string
	SearchUrl string
}

var SearchServices = []SearchService{
	SearchService{
		Domain:    "KatonSklep",
		SearchUrl: "https://katonsklep.pl/szukaj?controller=search&s=exia+repair",
	},
}
