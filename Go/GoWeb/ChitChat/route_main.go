package main

import "net/http"

// ハンドラ関数index---HTMLを生成してResponseWriterに書き出す
// ハンドラ関数は第1引数にResponseWriterをとり、第2引数にRequestへのポインタをとる
func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		// errがあれば一般用のナビゲーションバー、なければ会員用のバーを表示する
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
