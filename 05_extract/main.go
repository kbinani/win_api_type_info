package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	. "github.com/kbinani/win_api_type_info"
)

func main() {
	index := os.Args[1]
	f, err := os.Open(index)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	destDir := os.Args[2]

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	regRemoveRefTag := regexp.MustCompile(`<ref[^>]*>([^<]*)</ref>`)
	filtered := regRemoveRefTag.ReplaceAllFunc(bytes, func(match []byte) []byte {
		if m := regRemoveRefTag.FindSubmatch(match); len(m) > 0 {
			return m[1]
		} else {
			return match
		}
	})
	var o doxygen
	if err := xml.Unmarshal(filtered, &o); err != nil {
		panic(err)
	}

	var file *compounddef
	for _, comp := range o.CompounddefList {
		if comp.Kind == "file" {
			file = new(compounddef)
			*file = comp
			break
		}
	}
	if file == nil {
		panic(fmt.Errorf("Error: compounddef with 'file' kind not found"))
	}

	sections := make(map[string]sectiondef)
	for _, sec := range file.Sectiondef {
		sections[sec.Kind] = sec
	}

	typedefList := make(map[string]string)
	for _, t := range sections["typedef"].MemberdefList {
		if strings.Contains(t.Name, "(") || strings.Contains(t.Name, ")") || strings.Contains(t.Type, "(") || strings.Contains(t.Type, ")") {
			continue
		}
		typedefList[t.Name] = t.Type
	}
	SaveToJson(typedefList, filepath.Join(destDir, "typedef.json"))

	structs := make(map[string]Struct)
	for _, m := range o.CompounddefList {
		if m.Kind != "struct" {
			continue
		}
		var s Struct
		s.Fields = []Field{}

		for _, sec := range m.Sectiondef {
			if sec.Kind != "public-attrib" {
				continue
			}
			for _, member := range sec.MemberdefList {
				if strings.Contains(member.Name, "@") {
					continue
				}
				var f Field
				f.Name = member.Name
				f.Type = member.Type
				if member.Argsstring != "" {
					if strings.HasPrefix(member.Argsstring, "[") {
						f.Type = member.Argsstring + member.Type
					} else if strings.HasPrefix(member.Argsstring, ")(") && strings.HasSuffix(member.Type, "*") {
						// This field is function pointer
						f.Type = strings.TrimSuffix(member.Type, "*") + strings.TrimPrefix(member.Argsstring, ")(")
					} else {
						panic(fmt.Errorf("Name=%s; Type=%s; Argsstring=%s", member.Name, member.Type, member.Argsstring))
					}
				}
				s.Fields = append(s.Fields, f)
			}
		}

		structs[m.Compoundname] = s
	}
	SaveToJson(structs, filepath.Join(destDir, "struct.json"))

	enums := make(map[string]Enum)
	for _, m := range sections["enum"].MemberdefList {
		if strings.HasPrefix(m.Name, "@") {
			continue
		}
		var e Enum
		for _, v := range m.EnumValueList {
			var em EnumMember
			em.Name = v.Name
			e.Members = append(e.Members, em)
		}
		enums[m.Name] = e
	}
	SaveToJson(enums, filepath.Join(destDir, "enum.json"))
}
