package greet

import "strings"

//exported function starts with capital letter
//other functions are unexported and starts with small letter
func Hello(name string)string{
	clean := normalizeName(name)

	return 	"Hello, " + clean + "!"
}

func normalizeName(name string) string {
	n := strings.TrimSpace(name)
	if n == ""{
		return "Guest"
	}
	return strings.ToUpper(n)
}