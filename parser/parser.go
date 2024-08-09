package parser

import (
    "log"

	. "github.com/mantinhas/confed/scanner"
)

type Parser struct {
    source []Token
    M map[string]Token
    current int
}

func NewParser(sourceTokens []Token) *Parser {
    p := &Parser{
        source: sourceTokens,
        M: make(map[string]Token),
        current: 0,
    }

    return p
}

func (p *Parser) Parse() {
    for {
        if p.isAtEnd() {return}

        var token Token = p.advance()

        switch token.Type {
        case KEYWORD:
            p.consume(EQUALS, "keyword token not followed by '=' error")
            value := p.parseRightValue()

            p.M[string(token.Value)] = value
        default:
            log.Fatalf("unexpected token: %s at line %d", TK_TYPE_NAMES[token.Type], token.Line)
        }
    }
}

func (p *Parser) parseRightValue() (token Token) {
    if p.isAtEnd() { log.Fatalf("Unexpected end. Expected a value.") }
    
    token = p.advance()

    switch token.Type {
    case STRING, INT, FLOAT, BOOL :
        return token
    default:
        log.Fatalf("invalid value for keyword error: '%s' token not valid as a value at line %d.", token.Value, token.Line)
    }
    return
}

func (p *Parser) isAtEnd() bool {
    return p.current >= len(p.source)
}

func (p *Parser) peek() Token{
    return p.source[p.current]
}

func (p *Parser) advance() (token Token) {
    token = p.source[p.current]
    p.current++
    return token
}


func (p *Parser) consume(expected_type TK_TYPE, error_string string) {
    actual := p.peek()

    if actual.Type != expected_type {
        log.Fatalf("%s: Expected: %s\tActual: %s\tat line %d.", error_string, TK_TYPE_NAMES[expected_type], TK_TYPE_NAMES[actual.Type], actual.Line)
    }

    p.current++
}
