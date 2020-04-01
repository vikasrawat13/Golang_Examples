package pkgs

import "fmt"

// variable declarations
var x int = 50

// variable declaration
var (
	y    = 50
	z    = 100
	name = `vikas rawat`
	a    = 10.4
	m    = true
)

/* %T represent variable Type
   %d represent integer
   %f represent float
   %t represent boolean
   %s represent string
   %v represent default variable type format
   %b represent binary
   %c reprsent character code
*/

// variable declarations
const ax int = 1022

// variable declarations
var a1, a2 int = 20, 30
var d = 43
var abc string

//SampleFunc user defined function to print different variables
func SampleFunc() {

	// variable declaration
	msg := "hello world"
	display := fmt.Println

	if x > 10 {
		display("x is greater than 10")
	} else {
		fmt.Println("x is smaller than 10")
	}

	display(msg)
	conversion(&msg)
	display(msg)
	fmt.Printf("value of ax is %v", ax)
	fmt.Printf("\nvalue of a is %.2f", a) // handle decimal precision
	display("\nHello " + ", I'm Vikas Rawat")
	fmt.Printf("%c", 33)
}

func conversion(x *string) {
	*x = "HELLO WORLD.!"
	//return x
}

/*// reverse a string
func strReverse(x string)string{
	for
}*/
