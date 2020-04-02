package main

import (
	"fmt"
	p "pkgs"
	s "stringops"
)

func main() {
	fmt.Printf("%c", 95)
	//calling all functions exists under pkgs
	p.SampleFunc()
	p.TypesofLoops()
	p.IfElseFunc()
	fmt.Println("Reverse of Vikas is", p.StrRev("vikas"))
	fmt.Println(s.StrReverse("vikas"))
	fmt.Println(int('a'))

	fmt.Println("-----------------------------------------")
	x := ""
	fmt.Println(s.GetAlpha(x))
	fmt.Println(s.GetNumbers(x))
	fmt.Println(s.GetSpecialChar(x))
}
