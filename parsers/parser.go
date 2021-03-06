package parsers

import (
	"breeze-generator/core"
	"strings"
)

//parser names
const (
	Breeze = "breeze"
)

//schema file suffix
const (
	BreezeFileSuffix = ".breeze"
)

var (
	instances = map[string]core.Parser{
		Breeze: &BreezeParser{},
	}
)

//GetParser get Parser by name
func GetParser(name string) core.Parser {
	return instances[strings.ToLower(strings.TrimSpace(name))]
}

//Register : register a new parser
func Register(parser core.Parser) {
	instances[parser.Name()] = parser
}
