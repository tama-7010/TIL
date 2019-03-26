package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	// "Example ~"という文字列を読み込む
	reader := strings.NewReader("Example of io.SectionReader\n")
	// 14byte目から7byte取り出す
	sectionReader := io.NewSectionReader(reader, 14, 7)
	// sectionReaderを読み込み、標準出力に書き出す
	io.Copy(os.Stdout, sectionReader)
}
