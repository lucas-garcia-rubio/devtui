// Package cpf provides Brazilian CPF number generation.
package cpf

import (
	"fmt"
	"math/rand/v2"
)

// Generate returns a randomly generated, valid Brazilian CPF (11 digits).
//
// It generates 9 random base digits and then computes the 2 verification
// digits according to the official algorithm.
func Generate() string {
	base := make([]int, 9)
	for i := range base {
		base[i] = rand.IntN(10)
	}

	d1 := calcDigit(base, 10)
	d2 := calcDigit(append(base, d1), 11)

	return fmt.Sprintf(
		"%d%d%d%d%d%d%d%d%d%d%d",
		base[0], base[1], base[2], base[3], base[4], base[5], base[6], base[7], base[8],
		d1, d2,
	)
}

func calcDigit(digits []int, weightStart int) int {
	sum := 0
	weight := weightStart
	for _, d := range digits {
		sum += d * weight
		weight--
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}
