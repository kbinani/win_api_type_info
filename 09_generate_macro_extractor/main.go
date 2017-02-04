package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	regDefine *regexp.Regexp = nil
)

func init() {
	regDefine = regexp.MustCompile(`^\s*#\s*define\s+([A-Za-z0-9_]*).*$`)
}

func main() {
	includedFileList := os.Args[1]
	destFilePath := os.Args[2]

	fileList, err := readAllLines(includedFileList)
	if err != nil {
		panic(err)
	}
	allDefs := make(map[string]struct{})
	for _, file := range fileList {
		defs, err := extractDefinePreprocessors(file)
		if err != nil {
			panic(err)
		}
		for _, def := range defs {
			allDefs[def] = struct{}{}
		}
	}

	f, err := os.Create(destFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "#include <stdafx.h>\n")
	for def := range allDefs {
		fmt.Fprintf(f, "#ifdef %s\n", def)
		fmt.Fprintf(f, "____%s(%s)\n", def, def)
		fmt.Fprintf(f, "#endif\n")
	}
}

func extractDefinePreprocessors(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()
	defs := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		//fmt.Printf("%s\n", line)
		if match := regDefine.FindSubmatch([]byte(line)); len(match) > 0 {
			defs = append(defs, string(match[1]))
		}
	}
	return defs, nil
}

func readAllLines(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()
	lines := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, nil
}
