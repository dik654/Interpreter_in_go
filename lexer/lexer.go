package lexer

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
