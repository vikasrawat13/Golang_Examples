package pkgs

import "fmt"

//TypesofLoops is UDF
func TypesofLoops() {
	// example of a for loop
	fmt.Println("\n----------------------Table of 9 --------------------------")
	for i := 1; i <= 10; i++ {
		fmt.Printf("\nvalue of 9 * %d = %d", i, i*9)
	}

	// another way of for loop
	fmt.Println("\n--------------------Table of 13--------------------------")
	for i := 1; i <= 10; {
		fmt.Printf("\nvalue of 13 * %d = %d", i, i*13)
		i++
	}

	i := 1
	// while loop
	fmt.Println("\n---------------------Table of 2--------------------------")
	for i <= 10 {
		fmt.Printf("\nvalue of 2 * %d = %d", i, i*2)
		i++
	}

}

//ReverString function for reverse a string
func ReverString(txt string) string {
	var str string
	sl := len(txt)
	for sl >= 1 {
		str = str + txt[sl-1:sl]
		sl--
	}
	return str
}
