package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/kbinani/win_api_type_info"
)

func main() {
	dumpFilePath := os.Args[1]

	structJsonFilePath := os.Args[2]
	enumJsonFilePath := os.Args[3]

	structOutFilePath := os.Args[4]
	enumOutFilePath := os.Args[5]

	dump, err := os.Open(dumpFilePath)
	if err != nil {
		panic(err)
	}
	defer dump.Close()

	structs := make(map[string]Struct)
	LoadFromJson(&structs, structJsonFilePath)
	for name, theStruct := range structs {
		theStruct.ByteSize = -1
		for i := 0; i < len(theStruct.Fields); i++ {
			theStruct.Fields[i].BitOffset = -1
		}
		structs[name] = theStruct
	}

	enums := make(map[string]Enum)
	LoadFromJson(&enums, enumJsonFilePath)

	s := bufio.NewScanner(dump)
	for s.Scan() {
		line := s.Text()
		tokens := strings.Split(line, "\t")
		if len(tokens) == 0 {
			panic(fmt.Errorf("len(tokens)=%d; tokens=%v", len(tokens), tokens))
		}
		switch tokens[0] {
		case "size":
			name := tokens[1]
			size, err := strconv.Atoi(tokens[2])
			if err != nil {
				panic(err)
			}
			theStruct, ok := structs[name]
			if ok {
				theStruct.ByteSize = int32(size)
				structs[name] = theStruct
			}
		case "offset":
			typename := tokens[1]
			fieldname := tokens[2]
			offset, err := strconv.Atoi(tokens[3])
			if err != nil {
				panic(err)
			}
			theStruct, ok := structs[typename]
			if ok {
				for i := 0; i < len(theStruct.Fields); i++ {
					if theStruct.Fields[i].Name == fieldname {
						theStruct.Fields[i].BitOffset = int32(offset)
						break
					}
				}
				structs[typename] = theStruct
			}
		case "enum":
			name := tokens[1]
			member := tokens[2]
			value, err := strconv.Atoi(tokens[3])
			if err != nil {
				panic(err)
			}
			if theEnum, ok := enums[name]; ok {
				for i := 0; i < len(theEnum.Members); i++ {
					if theEnum.Members[i].Name == member {
						theEnum.Members[i].Value = int32(value)
						enums[name] = theEnum
						break
					}
				}
			}
		default:
			panic(fmt.Errorf("unknown token: %s", tokens[0]))
		}
	}
	SaveToJson(&enums, enumOutFilePath)

	filtered := make(map[string]Struct)
	for name, theStruct := range structs {
		if theStruct.ByteSize <= 0 {
			continue
		}
		reject := false
		for _, field := range theStruct.Fields {
			if field.BitOffset < 0 {
				reject = true
				break
			}
		}
		if reject {
			continue
		}
		filtered[name] = theStruct
	}
	SaveToJson(filtered, structOutFilePath)
}
