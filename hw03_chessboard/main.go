package main

import "fmt"

func main() {
	length := 8
	height := 8

	for i := 1; i <= height; i++ {
		fmt.Print("|")
		for j := 1; j <= length; j++ {
			switch {
			case i%2 == 1 && j%2 == 1:
				fmt.Print(" ")
			case i%2 == 1 && j%2 == 0:
				fmt.Print("#")
			case i%2 == 0 && j%2 == 1:
				fmt.Print("#")
			case i%2 == 0 && j%2 == 0:
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}
}
