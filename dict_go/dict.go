package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	workArray := [10]uint8{0}
	for indx := range workArray {
		fmt.Println(workArray[indx])
	}
}
