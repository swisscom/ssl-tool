package main

import "fmt"

func printTitle(title string) {
	fmt.Println(boldStyle.Render(title))
}
