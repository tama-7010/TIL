package main

import (
	"errors"
	"net/http"
)

// requestからクッキーを取り出す
func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, msg string) {
	cookie, err := request.Cookie("_cookie")
	// cookieが存在していれば、2番目のチェックにすすむ
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		// データベースを検索してこのセッションのUUIDが存在するかを調べる
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
