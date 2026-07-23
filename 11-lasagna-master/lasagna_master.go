package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutes int) int {
	if minutes == 0 {
		minutes = 2
	}
	return len(layers) * minutes
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodlesAmount int
	var sauceAmount int
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodlesAmount += 1
		}
		if layers[i] == "sauce" {
			sauceAmount += 1
		}
	}
	return noodlesAmount * 50, float64(sauceAmount) * 0.20
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	lastIngredient := friendList[len(friendList)-1]
	myList[len(myList)-1] = lastIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
	var onePortion []float64
	for i := 0; i <= len(quantities)-1; i++ {
		onePortion = append(onePortion, (quantities[i]*float64(portion))/2)
	}
	return onePortion
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
