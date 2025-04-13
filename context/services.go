package context

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
