package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var mass int

	f, err := os.Open("./input")
	if err != nil {
		log.Fatalln("Cant read file")
	}
	defer f.Close()

	var requiredFuel int
	s := bufio.NewScanner(f)
	for s.Scan() {
		_, err := fmt.Sscanf(s.Text(), "%d", &mass)
		if err != nil {
			log.Fatalln("Cant read line from file")
		}

		requiredFuel += calculateFuel(mass)
	}

	fmt.Printf("Required fuel: %d\n", requiredFuel)
}

func calculateFuel(mass int) int {
	return int(mass/3) - 2
}
