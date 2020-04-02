package stringops

import "regexp"

//GetSpecialChar function to extract special characters from a string
func GetSpecialChar(str string) string {
	reg, err := regexp.Compile("[a-zA-Z0-9]+")
	if err != nil {
		return "pattern not found"
	}
	return reg.ReplaceAllString(str, "")
}
