package stringops

import "regexp"

//GetNumbers function to extract numbers from a given string
func GetNumbers(str string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return "pattern not found"
	}
	return reg.ReplaceAllString(str, "")
}
