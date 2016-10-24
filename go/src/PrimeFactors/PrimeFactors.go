package PrimeFactors

func Generate(number int) []int {
	var factors []int

	if (number == 2) {
		factors = append(factors, 2)
	}

	return factors
}
