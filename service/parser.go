package service

type GundamOffers struct {
	Title  string
	Price  float64
	Link   string
	Domain string
}

func GetOffers() []*GundamOffers {
	var offers []*GundamOffers

	return offers
}
