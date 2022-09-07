package main

import "fmt"

// Maps

func main() {

	colors := map[string]string{
		"red":   "#ff000A",
		"green": "#00ff00",
	}

	colors["white"] = "#ffffff"

	delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex is", hex, ", color is", color)
	}
}
