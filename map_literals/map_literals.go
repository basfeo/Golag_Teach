package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	f := i % 100
	s := i % 100
	t := i % 10
	if f == s {
		fmt.Println("NO")
	} else if s == t {
		fmt.Println("NO")
	} else if f == t {
		fmt.Println("NO")
	}
	fmt.Println("YES")
}
