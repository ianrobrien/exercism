package lasagna

func PreparationTime(layers []string, averageTimePerLayer int) int {
	timePerLayer := averageTimePerLayer
	if averageTimePerLayer <= 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	amountNoodles := 0
	amountSauce := 0.0

	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			amountNoodles += 50
		case "sauce":
			amountSauce += .2
		}
	}

	return amountNoodles, amountSauce
}

func AddSecretIngredient(friendsRecipe []string, myRecipe []string) {
	secretIngredient := friendsRecipe[len(friendsRecipe)-1]
	myRecipe[len(myRecipe)-1] = secretIngredient
}

func ScaleRecipe(amountsNeeded []float64, portions int) []float64 {
	scaleFactor := float64(portions) / float64(2)
	var scaledAmounts []float64
	for i := 0; i < len(amountsNeeded); i++ {
		scaledAmounts = append(scaledAmounts, amountsNeeded[i]*float64(scaleFactor))
	}
	return scaledAmounts
}
