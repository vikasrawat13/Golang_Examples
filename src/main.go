package main

import (
	"fmt"
	p "pkgs"
)

func main() {
	//calling all functions exists under pkgs
	p.SampleFunc()
	p.TypesofLoops()
	p.IfElseFunc()
	fmt.Println("Reverse of Vikas is", p.ReverString(`vikas`))
}
