package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"log"
	"os"
	"regexp"
)

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

func absPathToRelative(lines iter.Seq[string], cwd string) iter.Seq[string] {
	pattern := regexp.MustCompile(`^(?P<prefix>\s*Error Trace:\s*)` + regexp.QuoteMeta(cwd) + `/(?P<path>[^:]+):(?P<line>[0-9]+)\s*$`)

	return func(yield func(string) bool) {
		for line := range lines {
			yield(pattern.ReplaceAllString(line, "${prefix}${path}:${line}"))
		}
	}
}

func reformat(lines iter.Seq[string]) iter.Seq[string] {
	return func(yield func(string) bool) {
		for line := range lines {
			for _, setting := range reformatSettings {
				match := setting.pattern.FindStringSubmatch(line)
				if len(match) != 0 {
					args := make([]any, len(match)-1)
					for i, m := range match[1:] {
						args[i] = m
					}
					line = fmt.Sprintf(setting.format, args...)

					break
				}
			}

			yield(line)
		}
	}
}

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

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}

	lines := readLines(os.Stdin)
	lines = absPathToRelative(lines, cwd)
	lines = reformat(lines)
	lines = colorize(lines)
	err = writeLines(os.Stdout, lines)
	if err != nil {
		log.Fatalf("failed to write output: %v", err)
	}
}
