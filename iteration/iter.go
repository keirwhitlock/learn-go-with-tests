package iteration

import "strings"

func Repeat(character string, count int, upper string) string {

	var repeated string
	for i := 0; i < count; i++ {
		if upper == "U" {
			character = strings.ToUpper(character)
		}
		repeated += character
	}
	return repeated
}
