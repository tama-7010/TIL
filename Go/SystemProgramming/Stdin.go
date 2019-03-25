package main

import (
	"io"
	"fmt"
	"os"
)

func main() {
	for {
		buffer := make([]byte, 5)
		// Go言語のRead()はブロックする
		// ノンブロッキングな処理はgoroutineを使用する
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input=%s\n", size, string(buffer))
	}
}