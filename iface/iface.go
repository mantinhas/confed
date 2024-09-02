package iface

import (
	"github.com/mantinhas/confed/parser"
	"github.com/mantinhas/confed/scanner"
)

func Get(key string, sourceString []byte) (string, bool) {
	s := scanner.NewScanner(sourceString)
	s.Scan()

	p := parser.NewParser(s.Tokens)
	p.Parse()

	value, ok := p.M[key]

	return string(value.Value), ok
}
