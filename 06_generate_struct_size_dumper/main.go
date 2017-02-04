package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	. "github.com/kbinani/win_api_type_info"
)

var (
	targetDir string
	blackList map[string]struct{}
)

func init() {
	blackList = make(map[string]struct{})
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func prepareBlacklist() {
	f, err := os.Open("blackList.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		m := s.Text()
		blackList[m] = struct{}{}
	}
}

func main() {
	prepareBlacklist()

	structJsonFilePath := os.Args[1]
	enumJsonFilePath := os.Args[2]
	targetDir = os.Args[3]

	for i := "A"[0]; i <= "Z"[0]; i++ {
		dir := filepath.Join(targetDir, string(i))
		os.Mkdir(dir, 0777)
	}
	os.Mkdir(filepath.Join(targetDir, "-"), 0777)
	os.Mkdir(filepath.Join(targetDir, "_"), 0777)

	structs := make(map[string]Struct)
	LoadFromJson(&structs, structJsonFilePath)

	enums := make(map[string]Enum)
	LoadFromJson(&enums, enumJsonFilePath)

	fmain, err := os.Create(filepath.Join(targetDir, "main.cpp"))
	if err != nil {
		panic(err)
	}
	defer fmain.Close()

	fmt.Fprintf(fmain, "#undef min\n")
	fmt.Fprintf(fmain, "#undef max\n")
	fmt.Fprintf(fmain, "#pragma warning(disable : 4995)\n")
	fmt.Fprintf(fmain, "#include <iostream>\n")
	fmt.Fprintf(fmain, "\n")

	fmt.Fprintf(fmain, "namespace size {\n")
	fmt.Fprintf(fmain, "namespace of {\n")
	for name, s := range structs {
		if shouldReject(name, s) {
			continue
		}
		fmt.Fprintf(fmain, "extern int %s();\n", name)
	}
	fmt.Fprintf(fmain, "} }\n")

	fmt.Fprintf(fmain, "namespace offset {\n")
	fmt.Fprintf(fmain, "namespace of {\n")
	for name, s := range structs {
		if shouldReject(name, s) {
			continue
		}
		fmt.Fprintf(fmain, "namespace %s {\n", name)
		for _, f := range s.Fields {
			fmt.Fprintf(fmain, "extern int %s();\n", f.Name)
		}
		fmt.Fprintf(fmain, "}\n")
	}
	fmt.Fprintf(fmain, "} }\n")

	fmt.Fprintf(fmain, "int main() {\n")

	var wg sync.WaitGroup

	for name, s := range structs {
		if shouldReject(name, s) {
			continue
		}

		printMain(fmain, name, s)

		wg.Add(1)
		go func(n string, theStruct Struct) {
			defer wg.Done()
			filename := fileNameForType(n)
			pre := string(filename[0])
			lines := printCFile(n, theStruct)
			filePath := filepath.Join(targetDir, pre, fmt.Sprintf("%s.cpp", filename))
			if err := printLinesIfChanged(lines, filePath); err != nil {
				panic(err)
			}
		}(name, s)
	}
	wg.Wait()

	for name, e := range enums {
		for _, em := range e.Members {
			fmt.Fprintf(fmain, "\tstd::cout << \"enum\\t%s\\t%s\\t\" << %s << std::endl;\n", name, em.Name, em.Name)
		}
	}

	fmt.Fprintf(fmain, "}\n")
}

func printLinesIfChanged(expectedLines []string, actualFileName string) error {
	shouldUpdate := false

	f, err := os.Open(actualFileName)
	if err != nil {
		shouldUpdate = true
	}

	if !shouldUpdate {
		s := bufio.NewScanner(f)
		i := 0
		for s.Scan() && i < len(expectedLines) {
			line := s.Text()
			if line != expectedLines[i] {
				shouldUpdate = true
				break
			}
			i++
		}
		if !shouldUpdate {
			if i == len(expectedLines) {
				return nil
			}
		}
	}
	f.Close()

	fw, err := os.Create(actualFileName)
	if err != nil {
		return err
	}
	defer fw.Close()
	for _, line := range expectedLines {
		fmt.Fprintf(fw, "%s\n", line)
	}
	return nil
}

func shouldReject(name string, theStruct Struct) bool {
	if strings.Contains(name, ":") {
		return true
	}
	if _, ok := blackList[name]; ok {
		return true
	}
	m := make(map[string]struct{})
	for _, f := range theStruct.Fields {
		if _, ok := m[f.Name]; ok {
			return true
		}
		m[f.Name] = struct{}{}
	}
	return false
}

func printMain(out io.Writer, typeName string, theStruct Struct) {
	fmt.Fprintf(out, "\tstd::cout << \"size\\t%s\\t\" << size::of::%s() << std::endl;\n", typeName, typeName)
	for _, field := range theStruct.Fields {
		fmt.Fprintf(out, "\tstd::cout << \"offset\\t%s\\t%s\\t\" << offset::of::%s::%s() << std::endl;\n", typeName, field.Name, typeName, field.Name)
	}
}

func printCFile(typeName string, theStruct Struct) []string {
	lines := []string{}
	lines = append(lines, fmt.Sprintf("#undef min"))
	lines = append(lines, fmt.Sprintf("#undef max"))
	lines = append(lines, fmt.Sprintf("#include <type_traits>"))
	lines = append(lines, fmt.Sprintf("#include \"../../../offsetof.hpp\""))
	lines = append(lines, fmt.Sprintf("namespace size {"))
	lines = append(lines, fmt.Sprintf("namespace of {"))
	lines = append(lines, fmt.Sprintf("class impl_%s {", typeName))
	lines = append(lines, fmt.Sprintf("\tstruct order2 {};"))
	lines = append(lines, fmt.Sprintf("\tstruct order1 : public order2 {};"))
	lines = append(lines, fmt.Sprintf(""))

	lines = append(lines, fmt.Sprintf("\ttemplate <typename>"))
	lines = append(lines, fmt.Sprintf("\tstruct enabler {"))
	lines = append(lines, fmt.Sprintf("\t\ttypedef bool type;"))
	lines = append(lines, fmt.Sprintf("\t};"))
	lines = append(lines, fmt.Sprintf(""))

	lines = append(lines, fmt.Sprintf("\ttemplate <typename = enabler<decltype(sizeof(::%s))>::type>", typeName))
	lines = append(lines, fmt.Sprintf("\tstatic int %s_impl(order1 arg0) {", typeName))
	lines = append(lines, fmt.Sprintf("\t\treturn sizeof(::%s);", typeName))
	lines = append(lines, fmt.Sprintf("\t}"))
	lines = append(lines, fmt.Sprintf(""))

	lines = append(lines, fmt.Sprintf("\tstatic int %s_impl(order2 arg0) {", typeName))
	lines = append(lines, fmt.Sprintf("\t\treturn -1;"))
	lines = append(lines, fmt.Sprintf("\t}"))
	lines = append(lines, fmt.Sprintf(""))

	lines = append(lines, fmt.Sprintf("public:"))
	lines = append(lines, fmt.Sprintf("\tstatic int %s() { return %s_impl(order1{}); }", typeName, typeName))
	lines = append(lines, fmt.Sprintf("};"))

	lines = append(lines, fmt.Sprintf("int %s() { return impl_%s::%s(); }", typeName, typeName, typeName))
	lines = append(lines, fmt.Sprintf("} }"))

	lines = append(lines, fmt.Sprintf("namespace offset {"))
	lines = append(lines, fmt.Sprintf("namespace of {"))
	lines = append(lines, fmt.Sprintf("namespace %s {", typeName))

	lines = append(lines, fmt.Sprintf("class impl {"))
	lines = append(lines, fmt.Sprintf("\tstruct order3 {};"))
	lines = append(lines, fmt.Sprintf("\tstruct order2 : public order3 {};"))
	lines = append(lines, fmt.Sprintf("\tstruct order1 : public order2 {};"))
	lines = append(lines, fmt.Sprintf(""))
	lines = append(lines, fmt.Sprintf("\ttemplate <typename>"))
	lines = append(lines, fmt.Sprintf("\tstruct enabler {"))
	lines = append(lines, fmt.Sprintf("\t\ttypedef bool type;"))
	lines = append(lines, fmt.Sprintf("\t};"))
	lines = append(lines, fmt.Sprintf(""))

	for _, field := range theStruct.Fields {
		lines = append(lines, fmt.Sprintf("\ttemplate <typename = std::enable_if<std::is_const<decltype(::%s::%s)>::value == false && std::is_array<decltype(::%s::%s)>::value == false>::type>", typeName, field.Name, typeName, field.Name))
		lines = append(lines, fmt.Sprintf("\tstatic int %s_impl(order1 arg0) {", field.Name))
		lines = append(lines, fmt.Sprintf("\t\treturn OFFSETBITOF(::%s, %s);", typeName, field.Name))
		lines = append(lines, fmt.Sprintf("\t}"))
		lines = append(lines, fmt.Sprintf(""))

		lines = append(lines, fmt.Sprintf("\ttemplate <typename = enabler<decltype(offsetof(::%s, %s))>::type>", typeName, field.Name))
		lines = append(lines, fmt.Sprintf("\tstatic int %s_impl(order2 arg0) {", field.Name))
		lines = append(lines, fmt.Sprintf("\t	return offsetof(::%s, %s) * 8;", typeName, field.Name))
		lines = append(lines, fmt.Sprintf("\t}"))
		lines = append(lines, fmt.Sprintf(""))

		lines = append(lines, fmt.Sprintf("\tstatic int %s_impl(order3 arg0) {", field.Name))
		lines = append(lines, fmt.Sprintf("\t	return -1;"))
		lines = append(lines, fmt.Sprintf("\t}"))
		lines = append(lines, fmt.Sprintf(""))
	}

	lines = append(lines, fmt.Sprintf("public:"))
	for _, field := range theStruct.Fields {
		lines = append(lines, fmt.Sprintf("\tstatic int %s() { return %s_impl(order1{}); }", field.Name, field.Name))
	}
	lines = append(lines, fmt.Sprintf("};"))

	for _, field := range theStruct.Fields {
		lines = append(lines, fmt.Sprintf("int %s() { return impl::%s(); }", field.Name, field.Name))
	}

	lines = append(lines, fmt.Sprintf("} } }"))

	return lines
}

func fileNameForType(typeName string) string {
	// Don't use typename directly as a filename to avoid file name conflict (ex: byte.go and BYTE.go)
	str := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(str); i++ {
		c := string(str[i])
		typeName = strings.Replace(typeName, c, "-"+strings.ToUpper(c), -1)
	}
	return typeName
}
