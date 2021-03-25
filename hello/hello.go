package main

import (
	"fmt"
)

const (
	Spanish = "Spanish"
	French = "French"
	HelloPrefix = "Hello, "
	SpanishHelloPrefix = "Hola, "
	FrenchHelloPrefix = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	var prefix string
	switch language {
	case Spanish:
		prefix = SpanishHelloPrefix
	case French:
		prefix = FrenchHelloPrefix
	default:
		prefix = HelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
