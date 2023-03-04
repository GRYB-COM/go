package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Powitanie(name string) (string, error) {
	if name == "" {
		return "", errors.New("brak imienia")
	}
	message := fmt.Sprintf(drawGreeting(), name)
	return message, nil
}
func drawGreeting() string {
	greetings := []string{
		"%v. Buźka chłopie",
		"Hejka %v",
		"%v witaj",
	}
	return greetings[rand.Intn(len(greetings))]
}
