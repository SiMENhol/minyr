package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/simenhol/minyr/yr"
)

func main() {
	// Venter for at brukeren bruker "minyr"
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter 'minyr' to start temperature conversion: ")
	text, _ := reader.ReadString('\n')
	if strings.ToLower(strings.TrimSpace(text)) != "minyr" {
		fmt.Println("Invalid input.")
		return
	}

	// Gir brukeren en liste for ting å velge
	fmt.Println("choose:")
	fmt.Println("  - 'convert' to convert temperature data from Celsius to Fahrenheit")
	fmt.Println("  - 'average' to get the average temperature for the entire period")
	fmt.Print("Enter convert or average: ")
	option, _ := reader.ReadString('\n')
	option = strings.ToLower(strings.TrimSpace(option))

	if option == "convert" {
		err := yr.convert()
		if err != nil {
			fmt.Println("Error during temperature conversion:", err)
			return
		}
		fmt.Println("Temperature conversion complete.")
		return
	}

	if option == "average" {
		fmt.Print("Enter unit of measurement ('c' for Celsius or 'f' for Fahrenheit): ")
		unit, _ := reader.ReadString('\n')
		unit = strings.ToLower(strings.TrimSpace(unit))

		avg, err := yr.Average(unit)
		if err != nil {
			fmt.Println("Error calculating average temperature:", err)
			return
		}
		fmt.Printf("Average temperature: %.2f %s\n", avg, unit)
	}

	// Wait for user input
	fmt.Println("Press enter to exit.")
	fmt.Scanln()
}
