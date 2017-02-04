package win_api_type_info

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"bytes"
	"io"
)

type Struct struct {
	Fields   []Field `json:"fields"`
	ByteSize int32   `json:"byte_size"`
}

type Field struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	BitOffset int32  `json:"bit_offset"`
}

type Enum struct {
	Members []EnumMember `json:"members"`
}

type EnumMember struct {
	Name  string `json:"name"`
	Value int32  `json:"value"`
}

func LoadFromJson(v interface{}, file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bytes, &v); err != nil {
		panic(err)
	}
}

func SaveToJson(v interface{}, file string) error {
	var encoded bytes.Buffer
	enc := json.NewEncoder(&encoded)
	enc.Encode(&v)

	var indented bytes.Buffer
	if err := json.Indent(&indented, encoded.Bytes(), "", "  "); err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	io.Copy(f, &indented)

	return nil
}
