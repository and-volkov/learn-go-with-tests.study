// Package hello
package hello

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	russian = "Russian"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	russianHelloPrefix = "Привет, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

// greetingPrefix returns the greeting prefix based on the language.
func greetingPrefix(language string) (prefix string) {
	// This is a named return variable `prefix`, so we can just assign to it
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case russian:
		prefix = russianHelloPrefix
	// if no matches before
	default:
		prefix = englishHelloPrefix
	}
	// This is a named return variable `prefix`, so we can just return it
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
