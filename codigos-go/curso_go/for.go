package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 1000000; {
		fmt.Println(i)
		i += 1
	}
}