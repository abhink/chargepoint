package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"cp/gol"
)

func main() {
	numGen := flag.Int("ng", 0, "number of generations to simulate")
	flag.Parse()
	grid := gol.NewGridFromString(readStdin())
	Simulate(grid, *numGen)
}

func Simulate(g *gol.Grid, numGen int) {
	fmt.Printf("%s\n", g)
	for i := 0; i < numGen; i++ {
		g = g.Advance()
		fmt.Printf("%s\n", g)
	}
}

func readStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	fmt.Println("Enter multiline text (Ctrl+D or Ctrl+Z on Windows to end):")

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading:", scanner.Err())
		return ""
	}

	return strings.Join(lines, "\n")
}
