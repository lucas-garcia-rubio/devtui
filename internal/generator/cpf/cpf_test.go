package cpf

import "testing"

func TestGenerate_ReturnsValidCPF(t *testing.T) {
	for range 100 {
		cpf := Generate()
		if len(cpf) != 11 {
			t.Fatalf("expected 11 digits, got %q (len=%d)", cpf, len(cpf))
		}

		digits := make([]int, 0, 11)
		for i := 0; i < len(cpf); i++ {
			c := cpf[i]
			if c < '0' || c > '9' {
				t.Fatalf("expected only digits, got %q", cpf)
			}
			digits = append(digits, int(c-'0'))
		}

		d1 := calcDigit(digits[:9], 10)
		if digits[9] != d1 {
			t.Fatalf("invalid first check digit for %q: got %d want %d", cpf, digits[9], d1)
		}
		d2 := calcDigit(digits[:10], 11)
		if digits[10] != d2 {
			t.Fatalf("invalid second check digit for %q: got %d want %d", cpf, digits[10], d2)
		}
	}
}
