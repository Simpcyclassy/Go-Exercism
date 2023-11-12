package lasagna

func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = 2
    }

    return len(layers) * time
}

func Quantities(layers []string) (noodles int, sauce float64) {
    for _, layer := range layers {
        if layer == "noodles" {
            noodles += 50
        } else if layer == "sauce" {
            sauce += 0.2
        }
    }

    return
}

func AddSecretIngredient (friendsList, myList []string) []string{
    myList[len(myList)-1] = friendsList[len(friendsList)-1]

    return myList
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i, quantity := range quantities {
		scaledQuantities[i] = quantity * float64(portions) / 2.0
	}

	return scaledQuantities
}
