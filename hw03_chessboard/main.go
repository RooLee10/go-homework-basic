package main

import (
	"fmt"
)

func main() {
	var length int
	fmt.Println("Укажите длину доски")
	_, err := fmt.Scanf("%d\n", &length)
	if err != nil {
		fmt.Println(err)
		return
	}

	var height int
	fmt.Println("Укажите высоту доски")
	_, err = fmt.Scanf("%d\n", &height)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= height; i++ {
		fmt.Print("|")
		for j := 1; j <= length; j++ {
			switch {
			case i%2 == 1 && j%2 == 1:
				fmt.Print("   ")
			case i%2 == 1 && j%2 == 0:
				fmt.Print("###")
			case i%2 == 0 && j%2 == 1:
				fmt.Print("###")
			case i%2 == 0 && j%2 == 0:
				fmt.Print("   ")
			}
		}
		fmt.Println("|")
	}
}
