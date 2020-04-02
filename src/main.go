package main

import (
	"fmt"
	t "topics"
	s "stringops"
)

func main() {
	var arrData [10]int

	fmt.Printf("%c", 95)
	//calling all functions exists under pkgs
	t.SampleFunc()
	t.TypesofLoops()
	t.IfElseFunc()
	fmt.Println("Reverse of Vikas is", t.StrRev("vikas"))
	fmt.Println(s.StrReverse("vikas"))
	fmt.Println(int('a'))

	fmt.Println("-----------------------------------------")
	x := "1AbT%34d1038dms#$@&^$"
	fmt.Println(s.GetAlpha(x))
	fmt.Println(s.GetNumbers(x))
	fmt.Println(s.GetSpecialChar(x))
	dArrays(&arrData)
	for _, value := range arrData {
		fmt.Println(value)
	}
}

func dArrays(arr *[10]int) {
	for i := 0; i < 10; i++ {
		(*arr)[i] = i
	}
}
