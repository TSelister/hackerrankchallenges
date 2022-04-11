package main

import "fmt"

func main() {
	var n, elem, sum int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&elem)
		sum += elem
	}

	fmt.Println(sum)
}
