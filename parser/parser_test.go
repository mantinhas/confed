package parser_test

import (
    "testing"

	"github.com/stretchr/testify/assert"
	"github.com/mantinhas/confed/scanner"
	"github.com/mantinhas/confed/parser"
)

func TestParserSimpleAttribution(t *testing.T) {
    result1 := scanner.Token{
            Type:  scanner.STRING,
            Value: []byte("john smith"),
            Line:  1,
        }
    result2 := scanner.Token{
            Type:  scanner.STRING,
            Value: []byte("usa"),
            Line:  2,
        }
    var sourceTokens []scanner.Token = []scanner.Token{
        {
            Type:  scanner.KEYWORD,
            Value: []byte("name"),
            Line:  1,
        },
        {
            Type:  scanner.EQUALS,
            Value: nil,
            Line:  1,
        },
        result1,
        {
            Type:  scanner.KEYWORD,
            Value: []byte("nationality"),
            Line:  2,
        },
        {
            Type:  scanner.EQUALS,
            Value: nil,
            Line:  2,
        },
        result2,
    }

    p := parser.NewParser(sourceTokens)

    p.Parse()

    assert.Equal(t, result1, p.M["name"])
    assert.Equal(t, result2, p.M["nationality"])
}
