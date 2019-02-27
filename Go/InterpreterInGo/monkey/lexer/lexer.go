// lexer/lexer.go

package lexer

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置（現在の文字を指し示す）
	readPosition int  // これから読み込む位置（現在の文字の次）
	ch           byte // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 終端チェック
		// 終端に達していた場合は l.ch を 0 にする
		l.ch = 0
	} else {
		// 終端に達していない場合は次の文字を l.ch にセットする
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
