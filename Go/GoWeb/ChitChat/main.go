package main

import (
	"html/template"
	"net/http"
)

func main() {
	// マルチプレクサを生成する
	mux := http.NewServeMux()
	// マルチプレクサは静的なファイルの返送にも使える
	files := http.FileServer(http.Dir("/public"))
	// 関数FileServerを使い、指定されたディレクトリからファイルを配信するハンドラを作成する
	// /static/以前を取り除き残った文字列が名前のファイルを探す
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// ルートURLをハンドラ関数にリダイレクトする
	// 第1引数にURLをとり、第2引数にハンドラ関数をとる
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", singup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
