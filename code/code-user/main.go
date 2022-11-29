package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scanln(&a, &b)
	_, _ = fmt.Println(a + b)
}
