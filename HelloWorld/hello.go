package main

import "fmt"

const danishHelloPrefix = "Hej, "

var helloPrefixes = map[string]string{
	"English": "Hello, ",
	"Spanish": "Hola, ",
	"French":  "Bonjour, ",
	"German":  "Hallo, ",
	"Italian": "Ciao, ",
}

func getHelloPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	case "German":
		prefix = "Hallo, "
	case "Italian":
		prefix = "Ciao, "
	default:
		prefix = "Hello, "
	}
	return
}

func HelloOld2() string {
	return "Hello, World!"
}
func HelloOld1(name string) string {
	return getHelloPrefix("English") + name
}
func HelloOld0(name string) string {
	return danishHelloPrefix + name
}
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = "English"
	}
	return helloPrefixes[language] + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
