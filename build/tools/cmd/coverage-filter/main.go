package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"iter"
	"log"
	"os"
	"strings"
)

func colorize(lines iter.Seq[string]) iter.Seq[string] {
	return func(yield func(string) bool) {
		for line := range lines {
			for _, setting := range styleSettings {
				line = setting.Apply(line)
			}

			yield(line)
		}
	}
}

func improveImportPath(lines iter.Seq[string], baseImportPath string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for line := range lines {
			yield(strings.ReplaceAll(line, baseImportPath, ""))
		}
	}
}

func readLines(source io.Reader) iter.Seq[string] {
	scanner := bufio.NewScanner(source)

	return func(yield func(string) bool) {
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				break
			}
		}
	}
}

func writeLines(target io.Writer, lines iter.Seq[string]) error {
	for line := range lines {
		_, err := fmt.Fprintln(target, line)
		if err != nil {
			return err
		}
	}

	return nil
}

func tabulate(lines iter.Seq[string]) iter.Seq[string] {
	return func(yield func(string) bool) {
		right := map[int]bool{2: true}
		widths := make([]int, 0)
		rows := make([][]string, 0)

		for line := range lines {
			fields := strings.Fields(line)
			if len(widths) < len(fields) {
				widths = append(widths, make([]int, len(fields)-len(widths))...)
			}
			for i, field := range fields {
				widths[i] = max(widths[i], len(field))
			}
			rows = append(rows, fields)
		}

		for i, width := range widths {
			base := log2i(width)
			base = max(3, base)
			l1 := 1 << base
			l0 := l1 >> 1
			if width < l0+l1 {
				widths[i] = l0 + l1
			} else {
				widths[i] = l1 + l1
			}
		}

		for _, row := range rows {
			var line strings.Builder
			for i, field := range row {
				if !right[i] {
					line.WriteString(field)
				}
				for j := len(field); j < widths[i]; j++ {
					line.WriteByte(' ')
				}
				if right[i] {
					line.WriteString(field)
				}
				if i < len(row)-1 {
					line.WriteByte(' ')
				}
			}
			yield(line.String())
		}
	}
}

func stripVerbose(lines iter.Seq[string]) iter.Seq[string] {
	return func(yield func(string) bool) {
		for line := range lines {
			if !verboseFilter.match(line) {
				yield(line)
			}
		}
	}
}

func main() {
	var verbose bool

	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Printf("Usage: %s [-v] <base-import-path>\n", os.Args[0])
		os.Exit(1)
	}

	baseImportPath := args[0] + "/"

	lines := readLines(os.Stdin)
	if !verbose {
		lines = stripVerbose(lines)
	}
	lines = improveImportPath(lines, baseImportPath)
	lines = tabulate(lines)
	lines = colorize(lines)
	err := writeLines(os.Stdout, lines)
	if err != nil {
		log.Fatal(err)
	}
}
