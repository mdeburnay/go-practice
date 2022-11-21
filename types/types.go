package types

import (
	"fmt"
)

var y = 42

var z string = "Shaken, not stirred"

func SayHey() {
	fmt.Println("Hey from types!")
	fmt.Println("%T\n", y)
	fmt.Println(z)
	fmt.Println("%T\n", z)
}
