package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	inFilePath := os.Args[1]
	outFilePath := os.Args[2]
	includedFilePath := os.Args[3]

	in, err := os.Open(inFilePath)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create(outFilePath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	regsRemove := []*regexp.Regexp{
		regexp.MustCompile(`__declspec\(deprecated\(".*"\)\)`),
		regexp.MustCompile(`__declspec\(uuid\(".*"\)\)`),
		regexp.MustCompile(`^\s*#pragma.*$`),
		regexp.MustCompile(`^\s*__pragma.*$`),
		regexp.MustCompile(`__declspec\(align\([0-9]*\)\)`),
	}

	replace := map[string]string{
		"__declspec(noreturn)": "",
		"__stdcall": "",
		"__cdecl": "",
		"__forceinline": "",
		"__inline": "",
		"__declspec(novtable)": "",
		"__declspec(dllimport)": "",
		"__declspec(noinline)": "",
		"__declspec(allocator)": "",
		"__declspec(deprecated)": "",
		"__declspec(restrict)": "",
		"__declspec(dllexport)": "",
		"__unaligned": "",
	}

	regsReject := []*regexp.Regexp{
		regexp.MustCompile(`^\s*$`),
	}

	regLinePragma := regexp.MustCompile(`^\s*#line\s+[0-9]+\s+\"(.*)\"\s*$`)

	files := []string{}
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := []byte(s.Text())
		if m := regLinePragma.FindSubmatch(line); len(m) > 0 {
			file := strings.ToLower(string(m[1]))
			found := false
			for _, f := range files {
				if f == file {
					found = true
					break
				}
			}
			if !found {
				files = append(files, file)
			}
		}

		for _, reg := range regsRemove {
			line = reg.ReplaceAllFunc(line, func(match []byte) []byte {
				return []byte{}
			})
		}

		for from, to := range replace {
			line = []byte(strings.Replace(string(line), from, to, -1))
		}

		reject := false
		for _, reg := range regsReject {
			if match := reg.FindSubmatch(line); len(match) > 0 {
				reject = true
				break
			}
		}
		if reject {
			continue
		}

		fmt.Fprintf(out, "%s\n", line)
	}

	f, err := os.Create(includedFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, file := range files {
		fmt.Fprintf(f, "%s\n", file)
	}
}
