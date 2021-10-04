package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layerPrepTime int) int {
	if layerPrepTime == 0 {
		layerPrepTime = 2
	}

	return len(layers) * layerPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(ingredients []string) (noodles int, sauce float64) {
	for _, ingredient := range ingredients {
		switch ingredient {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numberOfPeople int) (result []float64) {
	for _, value := range quantities {
		result = append(result, value/2.0*float64(numberOfPeople))
	}

	return result
}
