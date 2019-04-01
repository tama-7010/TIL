package main

import "fmt"

func main() {
	fmt.Println("start sub()")
	// 終了を受け取るためのチャンネル
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is Finished")
		// 終了を通知
		done <- true
	}()
	// 終了を待つ
	<-done
	fmt.Println("all tasks are finished")
}
