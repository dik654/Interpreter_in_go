package lexer

import (
	// *testsing.T를 가져와서 t.Fatal로 테스트 가능
	"testing"
	// token.go에 정의한 타입 가져오기
	"monkey/token"
)

type Lexer struct {
	input string
	// 현재 문자
	position int
	// 현재 다음 문자
	readPosition int
	// 현재 조사중인 문자
	ch byte
}

func New(input string) *Lexer {
	l := Lexer{input: input}
	return l
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	// 테스트할 타입과 예상하는 결과값을 갖는 구조체 생성
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	// TDD 테스트 주도 개발로, 만들어야할 함수보다 테스트코드를 먼저 작성하여 실제 함수를 나중에 작성해 테스트코드를 성공시키는 방식
	// New함수는 아직 작성하지 않았음
	// $ go test ./lexer
	// # monkey/lexer [monkey/lexer.test]
	// lexer\lexer_test.go:30:7: undefined: New
	// FAIL    monkey/lexer [build failed]
	// FAIL
	// Lexer를 반환하는 함수를 작성
	l := New(input)

	// test 구조체에 있는 모든 test case들 테스트
	// go test ./lexer로 테스트 실행가능
	for i, tt := range tests {
		// NextToken메서드는 아직 작성하지 않았음
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			//%q는 string에 ""을 붙여주는 서식 ex. "hello"
			t.Fatal("tests[%d]  - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
