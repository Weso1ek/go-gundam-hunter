package context

type GundamOffer struct {
	Title string
	//Price    float64
	Url string
	//Domain   string
	Img string
	//Active   bool
	//Discount *float64
}

type SearchService struct {
	Domain    string
	SearchUrl string
}

var SearchServices = []SearchService{
	SearchService{
		Domain:    "KatonSklep",
		SearchUrl: "https://katonsklep.pl/szukaj?controller=search&s=",
	},
}
