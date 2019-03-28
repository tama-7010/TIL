# Goroutineとは
goroutineはGoのランタイムに管理される軽量なスレッドのこと

```
go f(x, y, z)
```
と書けば、新しいgoroutineが実行される

```
f(x, y, z)
```
f, x, y, zの評価は、実行元（current）のgoroutineで実行され、fの実行は、新しいgoroutineで実行される

## Channelとは
チャネル（Channel）型は、チャネルオペレータの`<-`を用いて値の送受信ができる通り道のこと

```go
ch <- v // vをチャネルchへ送信する
v := <-ch // chから受信した変数をvへ割り当てる
```

マップとスライスのように、チャネルは使う前に以下のように生成する
```go
ch := make(chan int)
```
通常、片方が準備できるまで送受信はブロックされる。これにより、明確なロックや条件変数がなくても、goroutineの同期を可能にする



