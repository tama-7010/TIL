package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "sample.com:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: sample.com\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// ヘッダを表示する
	fmt.Println(res.Header)
	// ボディを表示して最後にClose()する
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
