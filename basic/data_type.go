package main

import (
	"fmt"
)

func main(){
	PrintFunc()
}
func PrintFunc() {
	var s1 = "Let's Go!"
	var i1 int32 = 18
	fmt.Printf("I'm %d years old.\n", i1)
	fmt.Println("Hello Gopher!")
	fmt.Print(s1)
}
