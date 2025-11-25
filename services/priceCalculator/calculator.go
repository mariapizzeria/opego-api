package priceCalculator

const (
	pet          = 50
	luggage_rack = 100
	english      = 200
	standard     = 100
	comfort      = 150
	comfort_plus = 200
	business     = 250
)

type Route struct {
	From string
	To   string
}

func PriceCalculation(tariff string, from string, to string, services []string) (int, error) {
	var total int
	switch tariff {
	case "standard":
		total = standard
	case "comfort":
		total = comfort
	case "comfort_plus":
		total = comfort_plus
	case "business":
		total = business
	default:
		total = standard
	}
	for _, service := range services {
		switch service {
		case "pet":
			total += pet
		case "luggage_rack":
			total += luggage_rack
		case "english":
			total += english
		default:
			total += 0
		}
	}

	route := &Route{
		From: from,
		To:   to,
	}

	routePrice, err := route.calculateRideCost(route)
	if err != nil {
		return 0, err
	}
	total += routePrice
	return total, nil
}

func (r *Route) calculateRideCost(route *Route) (int, error) {
	var total int
	// разобраться с api яндекс маршрутизатора
	// от предоставленного времени в дороге зависит цена
	return total, nil
}
