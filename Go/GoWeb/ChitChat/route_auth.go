package main

import (
	"net/http"
)

// POST /authenticate
// 与えられたユーザーとパスワードで認証をおこなう
func authenticate(writer http.ResponseWriter, request *http.Request) {
	r.ParseForm()

	// UserByEmail---メールアドレスを与えられると構造体Userを返す
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}

	// Encrypt---渡された文字列を暗号化する
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie", // 構造体の名前は任意。ブラウザに保存してほしいユニークなデータ
			Value:    session.Uuid,
			HttpOnly: true, // クッキーのアクセスをHTTPとHTTPSに限定する（JavaScriptなどは禁止される）
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}
