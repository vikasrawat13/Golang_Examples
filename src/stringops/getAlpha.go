package stringops

import "regexp"

//GetAlpha function to extract alphabets from a string
func GetAlpha(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		return "pattern not found"
	}
	return reg.ReplaceAllString(str, "")
}
