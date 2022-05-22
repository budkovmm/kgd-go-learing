package hello

import "fmt"

const (
	englishHelloPrefix = "Hello,"
	spanishHelloPrefix = "Holla,"
	Spanish            = "Spanish"
)

func getPrefixByLanguage(lang string) string {
	switch lang {
	case Spanish:
		return spanishHelloPrefix
	default:
		return englishHelloPrefix
	}
}

// Hello function generate greeting message
func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	helloPrefix := getPrefixByLanguage(lang)
	return fmt.Sprintf("%v %v", helloPrefix, name)
}

func main() {
	fmt.Println(Hello("Mikhail", ""))
}
