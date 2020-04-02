package stringops

//StrReverse function is used to revese a string
//Return the reversed string or 'empty string given' when an blank string passed
func StrReverse(str string) string {
	var txt string
	if len(str) == 0 {
		return "empty string given"
	}
	for i := len(str); i > 0; i-- {
		txt = txt + string(str[i-1])
	}
	return txt
}
