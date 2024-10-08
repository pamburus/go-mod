package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	tracePattern := regexp.MustCompile(`^(?P<prefix>\s*Error Trace:\s*)` + regexp.QuoteMeta(cwd) + `/(?P<path>[^:]+):(?P<line>[0-9]+)\s*$`)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = tracePattern.ReplaceAllString(line, "${prefix}${path}:${line}")

		for _, setting := range styleSettings {
			line = setting.Apply(line)
		}

		fmt.Println(line)
	}
}
