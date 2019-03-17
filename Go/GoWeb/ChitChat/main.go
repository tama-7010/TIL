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

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

// ハンドラ関数index---HTMLを生成してResponseWriterに書き出す
// ハンドラ関数は第1引数にResponseWriterをとり、第2引数にRequestへのポインタをとる
func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
