package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// 렉서가 토큰, 문자를 알 수 없다
	ILLEGAL = "ILLEGAL"
	// 파일의 끝, 렉서 멈춤
	EOF = "EOF"

	// 변수 식별자
	IDENT = "IDENT"
	// 타입
	INT = "INT"
	// 연산자
	ASSIGN = "="
	PLUS   = "+"
	// 구분자
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	//예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
