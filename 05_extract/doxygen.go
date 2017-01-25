package main

type enumvalue struct {
	Name string `xml:"name"`
}

type memberdef struct {
	Kind          string      `xml:"kind,attr"`
	Name          string      `xml:"name"`
	Type          string      `xml:"type"`
	Argsstring    string      `xml:"argsstring"`
	EnumValueList []enumvalue `xml:"enumvalue"`
}

type sectiondef struct {
	Kind          string      `xml:"kind,attr"`
	MemberdefList []memberdef `xml:"memberdef"`
}

type compounddef struct {
	Compoundname string       `xml:"compoundname"`
	Kind         string       `xml:"kind,attr"`
	Sectiondef   []sectiondef `xml:"sectiondef"`
}

type doxygen struct {
	CompounddefList []compounddef `xml:"compounddef"`
}
