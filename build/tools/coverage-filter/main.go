package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: coverage-filter <base-import-path>")
		os.Exit(1)
	}

	baseImportPath := os.Args[1] + "/"

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, baseImportPath, "")

		for _, setting := range styleSettings {
			line = setting.Apply(line)
		}

		fmt.Println(line)
	}
}
