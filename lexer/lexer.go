package lexer

import "monkey/token"

type Lexer struct {
	input string
	// 현재 문자 index
	position int
	// 다음 문자 index
	readPosition int
	// 현재 조사중인 문자
	ch byte
}

// lexer 구조체를 생성해서 인수로 들어온 string을 input으로 설정
// lexer 구조체 포인터를 리턴
func New(input string) *Lexer {
	l := &Lexer{input: input}
	// 첫 번째 문자를 ch에 읽어두기 위해 readChar메서드 실행
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// readPosition이 파일 전체 길이를 넘어섰다면
	if l.readPosition >= len(l.input) {
		// EOF라고 표기
		l.ch = 0
	} else {
		// ch에 현재 문자 추가
		l.ch = l.input[l.readPosition]
	}
	// 이후 readChar을 다시 실행했을 때를 위해 업데이트
	// 현재 문자 index 업데이트
	l.position = l.readPosition
	// 다음 문자 index 업데이트
	l.readPosition += 1
}

// ch에 들어있는 문자에 해당하는 토큰 생성
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	// 다음 문자를 위한 준비
	l.readChar()
	return tok
}

// NextToken메서드에서 token을 생성하는데 사용하는 함수
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// $ go test ./lexer
// ok      monkey/lexer    0.173s
