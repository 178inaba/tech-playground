package checkdigit

import "errors"

// A Generator generates a check digit by implemented algorithm or calculator.
type Generator interface {
	Generate(seed string) (int, error)
}

// A Verifier is verifying to code by implemented algorithm or calculator.
type Verifier interface {
	Verify(code string) bool
}

type Provider interface {
	Verifier
	Generator
}

type luhn struct{}

func NewLuhn() Provider {
	return &luhn{}
}

// Verify implements checkdigit.Verifier interface.
func (l luhn) Verify(code string) bool {
	if len(code) < 2 {
		return false
	}
	i, err := l.Generate(code[:len(code)-1])

	return err == nil && i == int(code[len(code)-1]-'0')
}

// Generate implements checkdigit.Generator interface.
func (l luhn) Generate(seed string) (int, error) {
	if seed == "" {
		return 0, ErrInvalidArgument
	}

	sum, parity := 0, (len(seed)+1)%2
	for i, n := range seed {
		if isNotNumber(n) {
			return 0, ErrInvalidArgument
		}
		d := int(n - '0')
		if i%2 == parity {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}

	return sum * 9 % 10, nil
}

type gtin struct {
	digit   int
	posCorr bool
}

// Verify implements checkdigit.Verifier interface.
func (g *gtin) Verify(code string) bool {
	if len(code) != g.digit {
		return false
	}
	i, err := g.Generate(code[:len(code)-1])

	return err == nil && i == int(code[len(code)-1]-'0')
}

// Generate implements checkdigit.Generator interface.
func (g *gtin) Generate(seed string) (int, error) {
	if len(seed) != g.digit-1 {
		return 0, ErrInvalidArgument
	}

	var oddSum, evenSum int
	for i, n := range seed {
		if isNotNumber(n) {
			return 0, ErrInvalidArgument
		}
		if g.posCorr {
			i++
		}
		if i%2 == 0 {
			evenSum += int(n - '0')
		} else {
			oddSum += int(n - '0')
		}
	}

	d := 10 - (evenSum*3+oddSum)%10
	if d == 10 {
		d = 0
	}

	return d, nil
}

// ErrInvalidArgument is happening when given the wrong argument.
var ErrInvalidArgument = errors.New("checkdigit: invalid argument")

func isNotNumber(n rune) bool {
	return n < '0' || '9' < n
}
