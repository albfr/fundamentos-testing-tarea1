package main

import (
	"example/user/hello/utils"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("I need your name -.-")
		os.Exit(1)
	}

	name := os.Args[1]
	fmt.Println(utils.Hello(name))

	var input string
	fmt.Scanln(&input)

	for input != "Stop!" {
		fmt.Println(utils.ReverseString(input))
		if utils.IsPalindome(input) {
			fmt.Println("¡Bonita Palabra!")
		}
		fmt.Scanln(&input)
	}

	fmt.Println("Adios", name)

}
