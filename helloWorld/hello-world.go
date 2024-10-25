package helloWorld

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishGreeting = "Hello, "
	spanishGreeting = "Hola, "
	frenchGreeting  = "Bonjour, "
)

func main() {
	fmt.Println(Hello("", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishGreeting
	case french:
		prefix = frenchGreeting
	default:
		prefix = englishGreeting
	}

	return
}
