package main

import (
	"example/user/hello/utils"
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	fmt.Println("%s", utils.Hello(name))

	var input string
	fmt.Scanln(&input)

	for input != "Stop!" {
		fmt.Println("%s", utils.ReverseString(input))
		if utils.IsPalindome(input) {
			fmt.Println("¡Bonita Palabra!")
		}
		fmt.Scanln(&input)
	}

	fmt.Println("Adios", name)

}
