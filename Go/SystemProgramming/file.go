package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("file.go")
	if err != nil {
		panic(err)
	}
	// defer 現在のスコープが終了したら、その後ろの行の処理を実行する
	defer file.Close()
	io.Copy(os.Stdout, file)
}
