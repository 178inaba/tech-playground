package checkdigit

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

// Verify implements checkdigit.Verifier interface.
func (l luhn) Verify(code string) bool {
	// TODO
	return false
}
