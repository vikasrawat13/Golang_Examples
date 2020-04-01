package pkgs

import "fmt"

//IfElseFunc is sample function for if else statements
func IfElseFunc() {
	a, b, c := 10, 12, 21
	print := fmt.Println

	print("\n----------------If condition example with logical operators -----------------")
	if a == 10 {
		print("Value of A is equal to 10")
	}
	if b >= 10 {
		print("Value of B is greater than 10")
	}
	if c > 10 && c <= 13 {
		print("Value of C is greater than 10 but less than 14")
	} else {
		print("Value is greater than 20")
	}

}
