package PrimeFactors

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	if Generate(1) != nil {
		t.Error("Expected []")
	}

	if Generate(2)[0] != 2 {
		t.Error("Expected [2]")
	}
}
