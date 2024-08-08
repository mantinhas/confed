package scanner

import (
	"log"
	"unicode"
)

type TK_TYPE int

const (
	STRING TK_TYPE = iota
	INT
	EQUALS
	KEYWORD
	BOOL

	LEFT_BRACKET
	RIGHT_BRACKET
	LEFT_BRACE
	RIGHT_BRACE
)

type Token struct {
	Type  TK_TYPE
	Value []byte
	Line  int
}

type Scanner struct {
	source  []byte
	start   int
	current int
	line    int
	Tokens  []Token
}

//type Config struct {
//	IOSource io.Reader
//}

func New(sourceString []byte) *Scanner {

	//var stringFile []byte
	//var err error

	//if cfg.IOSource != nil {
	//	stringFile, err = io.ReadAll(bufio.NewReader(cfg.IOSource))
	//} else {
	//	stringFile, err = io.ReadAll(bufio.NewReader(os.Stdin))
	//}

	//if err != nil {
	//	log.Fatal(err)
	//}

	s := &Scanner{
		start:   0,
		current: 0,
		line:    1,
		source:  sourceString,
		Tokens:  []Token{},
	}

	return s
}

func (s *Scanner) Scan() {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
}

func (s *Scanner) scanToken() {
	value := s.advance()

	switch {
	case value == '=':
		s.Tokens = append(s.Tokens, Token{
			Type:  EQUALS,
			Value: nil,
			Line:  s.line,
		})
	case value == '"':
		s.scanString()
	case value == '\n':
		s.line += 1
	case value == ' ':
		return
	case value == 0:
		return
		//s.Tokens = append(s.Tokens, Token{
		//	Type:	EOF,
		//	value:	nil,
		//})
	case isValidCharForBareKey(value):
		s.scanKeyword()
	}
}

func isValidCharForBareKey(value byte) bool {
	return unicode.IsLetter(rune(value)) || value == '_' || value == '-' || unicode.IsDigit(rune(value))
}

func (s *Scanner) scanKeyword() {
	for isValidCharForBareKey(s.peek()) {
		s.advance()
	}

	keyword := s.source[s.start:s.current]

	switch string(keyword) {
	case "true":
		s.Tokens = append(s.Tokens, Token{
			Type:  BOOL,
			Value: keyword,
			Line:  s.line,
		})
	case "false":
		s.Tokens = append(s.Tokens, Token{
			Type:  BOOL,
			Value: keyword,
			Line:  s.line,
		})
	default:
		s.Tokens = append(s.Tokens, Token{
			Type:  KEYWORD,
			Value: keyword,
			Line:  s.line,
		})
	}
}

func (s *Scanner) scanString() {
	for s.peek() != '"' && !s.isAtEnd() {
		s.advance()
	}
	if s.isAtEnd() {
		log.Fatalf("Unterminated string at line %d:\t %s", s.line, s.source[s.start:])
	}
	// Consume closing quote
	s.advance()

	s.Tokens = append(s.Tokens, Token{
		Type:  STRING,
		Value: s.source[s.start+1 : s.current-1],
		Line:  s.line,
	})
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() (value byte) {
	value = s.source[s.current]
	s.current += 1
	return value
}

func (s *Scanner) peek() (value byte) {
	if s.isAtEnd() {
		return 0
	}
	value = s.source[s.current]
	return value
}

func (s *Scanner) peekNext() (value byte) {
	if s.current+1 >= len(s.source) {
		return 0
	}
	value = s.source[s.current+1]
	return value
}
