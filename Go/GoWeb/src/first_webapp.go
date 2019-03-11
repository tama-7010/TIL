package main

import (
	"fmt"
	"net/http"
	// 「net/http」---実行環境を提供し、コード共にコンパイルされるので、
	// そのまま配置可能なスタンドアローンのWebアプリが生成される。
)

func handler(writer http.ResponseWriter, request *http.Request) {
	/* 「ハンドラ（handler）」という言葉は、イベントがトリガー（きっかけ）として起動されるコールバック関数でよく使われる。*/
	/* handler関数には2つの引数がある。インターフェースResponseWriteと構造体Requestへのポインタ。*/
	/* この関数がRequestから情報を取り出してHTTPレスポンスを生成し、そのレスポンスがResponseWriterを介して通信される。*/
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:]) // パスの:8080以降が追加される
}

func main() {
	http.HandleFunc("/", handler)     // ルートURL(/)が呼び出されたとき、上で定義した「handler」が起動されるように設定する。
	http.ListenAndServe(":8080", nil) //　8080ポートを監視するようにサーバを起動する。
}
