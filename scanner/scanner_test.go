package scanner_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mantinhas/confed/scanner"
	"github.com/mantinhas/confed/utils"
)

func TestScannerStringAttributionTokensTypes(t *testing.T) {
	testdata := utils.ReadFile("testdata/string_attribution_test.toml")
	s := scanner.NewScanner(testdata)

	assert.Equal(t, 0, len(s.Tokens))
	s.Scan()
	assert.Equal(t, 6, len(s.Tokens))

	var expectedTK_TYPE []scanner.TK_TYPE = []scanner.TK_TYPE{
		scanner.KEYWORD,
		scanner.EQUALS,
		scanner.STRING,
		scanner.KEYWORD,
		scanner.EQUALS,
		scanner.STRING,
	}

	assert.Equal(t, expectedTK_TYPE, getTokensType(s.Tokens))
}

func TestScannerStringAttributionTokenValues(t *testing.T) {
	testdata := utils.ReadFile("testdata/string_attribution_test.toml")
	s := scanner.NewScanner(testdata)

	s.Scan()

	assert.Equal(t, 6, len(s.Tokens))

	assert.Equal(t, scanner.Token{
		Type:  scanner.KEYWORD,
		Value: []byte("name"),
		Line:  1,
	}, s.Tokens[0])
	assert.Equal(t, scanner.Token{
		Type:  scanner.EQUALS,
		Value: nil,
		Line:  1,
	}, s.Tokens[1])
	assert.Equal(t, scanner.Token{
		Type:  scanner.STRING,
		Value: []byte("john smith"),
		Line:  1,
	}, s.Tokens[2])
	assert.Equal(t, scanner.Token{
		Type:  scanner.KEYWORD,
		Value: []byte("nationality"),
		Line:  2,
	}, s.Tokens[3])
	assert.Equal(t, scanner.Token{
		Type:  scanner.EQUALS,
		Value: nil,
		Line:  2,
	}, s.Tokens[4])
	assert.Equal(t, scanner.Token{
		Type:  scanner.STRING,
		Value: []byte("usa"),
		Line:  2,
	}, s.Tokens[5])
}

func TestScannerBoolAttributionTokenValues(t *testing.T) {
	testdata := utils.ReadFile("testdata/bool_attribution_test.toml")
	s := scanner.NewScanner(testdata)

	s.Scan()

	assert.Equal(t, 6, len(s.Tokens))

	assert.Equal(t, scanner.Token{
		Type:  scanner.KEYWORD,
		Value: []byte("isUp_"),
		Line:  1,
	}, s.Tokens[0])
	assert.Equal(t, scanner.Token{
		Type:  scanner.EQUALS,
		Value: nil,
		Line:  1,
	}, s.Tokens[1])
	assert.Equal(t, scanner.Token{
		Type:  scanner.BOOL,
		Value: []byte("true"),
		Line:  1,
	}, s.Tokens[2])
	assert.Equal(t, scanner.Token{
		Type:  scanner.KEYWORD,
		Value: []byte("ISDOWN"),
		Line:  3,
	}, s.Tokens[3])
	assert.Equal(t, scanner.Token{
		Type:  scanner.EQUALS,
		Value: nil,
		Line:  3,
	}, s.Tokens[4])
	assert.Equal(t, scanner.Token{
		Type:  scanner.BOOL,
		Value: []byte("false"),
		Line:  3,
	}, s.Tokens[5])
}

func TestScannerNumberAttributionTokenValues(t *testing.T) {
	testdata := utils.ReadFile("testdata/number_attribution_test.toml")
	s := scanner.NewScanner(testdata)
	s.Scan()

	assert.Equal(t, 6, len(s.Tokens))

	assert.Equal(t, scanner.Token{
		Type:  scanner.KEYWORD,
		Value: []byte("age"),
		Line:  1,
	}, s.Tokens[0])
	assert.Equal(t, scanner.Token{
		Type:  scanner.EQUALS,
		Value: nil,
		Line:  1,
	}, s.Tokens[1])
	assert.Equal(t, scanner.Token{
		Type:  scanner.INT,
		Value: []byte("16"),
		Line:  1,
	}, s.Tokens[2])
	assert.Equal(t, scanner.Token{
		Type:  scanner.KEYWORD,
		Value: []byte("grade"),
		Line:  2,
	}, s.Tokens[3])
	assert.Equal(t, scanner.Token{
		Type:  scanner.EQUALS,
		Value: nil,
		Line:  2,
	}, s.Tokens[4])
	assert.Equal(t, scanner.Token{
		Type:  scanner.FLOAT,
		Value: []byte("20.170"),
		Line:  2,
	}, s.Tokens[5])
}

func getTokensType(tokens []scanner.Token) []scanner.TK_TYPE {
	var types []scanner.TK_TYPE
	for _, token := range tokens {
		types = append(types, token.Type)
	}
	return types
}
