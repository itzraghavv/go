package main

import (
	"fmt"

	l2custompack "github.com/itzraghavv/custompack"
)

func test(text string) {
	fmt.Println(text)
}

func main() {
	test("starting Mailio server")
	fmt.Println(l2custompack.Reverse("hello"))
}
