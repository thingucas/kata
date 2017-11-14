package PrimeFactors

type PrimeFactorGenerator interface {
	Generate(number int) []int
}

type PrimeFactors struct {
}

func (p *PrimeFactors) Generate(number int) (factors []int) {
	return
}
