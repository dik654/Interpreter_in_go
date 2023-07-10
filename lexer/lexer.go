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
		// ch에 다음 문자 넣어두기
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
	// monkey 언어에서 공백은 의미가 없으므로 넘어가는 로직 추가
	l.skipWhitespace()

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
	// 글자인지 판별하는 과정 추가
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			// 글자 판별중에 예약어인지 체크하는 부분 추가
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
			// 숫자라면
		} else if isDigit(l.ch) {
			// 타입을 INT로 설정하고
			tok.Type = token.INT
			// 이어진 숫자 다 읽기 ex 100이면 100 그대로
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	// 다음 문자를 위한 준비
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	// 공백이거나 공백문자 등의 경우
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		// 다음 문자 읽기(공백을 무시하기 위해)
		l.readChar()
	}
}

// 식별자를 읽은 뒤 글자가 아닐때까지 읽어들인다. ex. let이면 l을 읽으면 t까지
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		// letter면 다음 문자 읽기
		l.readChar()
	}
	return l.input[position:l.position]
}

// 글자인지 검사
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		// 숫자라면 계속 읽어들여서
		l.readChar()
	}
	// 전체 숫자를 리턴
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	// 숫자인지 체크
	return '0' <= ch && ch <= '9'
}

// NextToken메서드에서 token을 생성하는데 사용하는 함수
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
