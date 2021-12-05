package matrix

import (
	"fmt"
)

var m [][]string = [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}

func ExampleEcho() {
	fmt.Println(echo(m))
    // Output: 
	//1,2,3
	//4,5,6
	//7,8,9
}

func ExampleInvert() {
	fmt.Println(invert(m))
    // Output: 
	//1,4,7
	//2,5,8
	//3,6,9
}

func ExampleFlatten() {
	fmt.Println(flatten(m))
    // Output: 
	//1,2,3,4,5,6,7,8,9
}

func ExampleSum() {
	fmt.Println(sum(m))
    // Output: 
	//45
}

func ExampleMultiply() {
	fmt.Println(multiply(m))
    // Output: 
	//362880
}

