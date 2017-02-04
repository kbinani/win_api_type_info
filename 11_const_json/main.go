package main

import (
	"os"
	"bufio"
	"regexp"

	. "github.com/kbinani/win_api_type_info"
)

func main() {
	cfilePath := os.Args[1]
	destPath := os.Args[2]

	cfile, err := os.Open(cfilePath)
	if err != nil {
		panic(err)
	}
	defer cfile.Close()

	regPP := regexp.MustCompile(`^____([a-zA-Z0-9_]*)\((.*)\)$`)

	consts := make(map[string]string)

	s := bufio.NewScanner(cfile)
	for s.Scan() {
		line := s.Text()
		if match := regPP.FindSubmatch([]byte(line)); len(match) > 0 {
			name := string(match[1])
			if name == "__FILEW__" { // may contains absolute file path
				continue
			}
			value := string(match[2])
			consts[name] = value
		}
	}

	SaveToJson(&consts, destPath)
}
