# HTTP入門
- HTTPはステートレスでテキストベースのリクエスト/レスポンス型プロトコルであり、クライアント/サーバ型のコンピューティングモデルで使用される。  
- 「リクエスト/レスポンス」は2台のコンピュータがお互いに「会話」をする際の基本的な方法。
- 「クライアント/サーバ型」のコンピュータモデルでは、常にリクエスト側がレスポンス側との会話を始める。

## HTTPリクエスト
HTTPはリクエストから始まる。  
HTTPリクエストは次の行で構成されている

1. リクエスト行---リクエストメソッド、URI、HTTPのバージョンの順に続く
2. リクエストヘッダ（0行以上）
3. 空行
4. メッセージ本体（省略可）

### HTTPリクエストの例
```
GET /Protocols/rfc2616/rfc2616.html HTTP/1.1
Host: www.w3.org
User-Agent: Mozilla/5.0
<<<空行>>>
```

### リクエストメソッド
リクエストメソッドはリクエスト行の最初の単語であり、リソースに対して行うべき処理を表す。

- GET---指定されたリソースを返す
- HEAD---メッセージ本体を返さない
- POST---メッセージ本体のデータをURIで示されたリソースに渡す。
- PUT---メッセージ本体のデータをURIで指定されたリソースにする。
- DELETE---指定されたリソースを削除する。
- TRACE---リクエストを送り返す。
- OPTIONS---サーバがサポートするHTTPメソッドのリストを返す。
- CONNECT---クライアントとのネットワーク接続を準備するようサーバに指示する。
- PATCH---メッセージ本体のデータに従ってリソースを変更する。

### リクエストヘッダ
リクエストヘッダはリクエストに付随する情報やクライアントに関するその他の情報がある。  
リクエストヘッダは「<名前>:<値>」の形式を持つプレーンテキストであり、CR+LFで終わる。

- Accept---クライアントがHTTPレスポンスの一部として受け取ることができるコンテンツタイプ
- Accept-Charset---サーバから受け取る際に要求される文字セット
- Authorization---ベーシック認証の証明書をサーバに送付する
- Cookie---サーバがそれまでにブラウザに対して2つのクッキーを設定したとすれば、Cookieには2つのクッキーが含まれており、それぞれが「;」で区切った名前と値の組で構成される。
- Content-Length---リクエスト本体の長さ（バイト数）
- Content-Type---リクエスト本体のコンテンツタイプ。
    - POSTかPUT : x-www-form-urlencoded
    - ファイルをアップロードする : multipart/form-data
- Host---サーバ名とポート番号
- Referrer---リクエストされたページにリンクしていた以前のページのアドレス
- User-Agent---呼び出しているクライアントに関する情報

## HTTPレスポンス
リクエストがあるたびにHTTPレスポンスのメッセージが1つ返される。  
HTTPリクエストと同様にHTTPレスポンスも何行かのプレーンテキストから成る。

1. ステータス行
2. 0個以上のレスポンスヘッダ
3. 空行
4. メッセージ本体（省略可能）

### HTTPレスポンスの例
```
200 OK
Date: Sat, 26 Nov 2016 12:58:58 GMT
Server: Apache/2
Last-Modified: Thu, 25 Aug 2016 21:01:33 GMT
Content-Length: 33115
Content-Type: text/html; charset=iso-8859-1

<!DOCTYPE html><html><head><title>HTTP/1.1</title></head><body>...</body></html>
```

### レスポンスのステータスコード
- 1XX---情報提供
- 2XX---成功
- 3XX---リダイレクション
- 4XX---クライアントエラー
- 5XX---サーバエラー

### レスポンスヘッダ
- Allow---どのリクエストメソッドをサポートしているかを伝える。
- Content-Length---レスポンス本体の長さ（バイト数）。
- Content-Type---レスポンス本体のコンテンツタイプ。
- Date---現在日時
- Location---リダイレクション時にURLをどこに要求すれば良いかを伝える。
- Server---レスポンスを返しているサーバのドメイン名。
- Set-Cockie---クライアントへのクッキーを設定できる。
- WWW-Authenticate---リクエストヘッダのAuthorizationフィールドにどのような認証情報を含めるべきかをクライアントに伝える。

### URI
URIの一般型
`＜スキーム名＞:＜階層部＞[?<クエリ＞][#<フラグメント>]`
- スキーム名---URIの構造の残りの部分を定義するURIスキームの名前。
- 階層部---リソースを特定する情報。
- クエリ---階層表示になじまない、その他の情報。
- フラグメント---定義されたURIの一部となる、2次リソースを特定するもの。

## Webアプリの構成
これまでをまとめると、Webアプリケーションとは以下のことをするプログラムだと言える。
 1. HTTPを介し、HTTPリクエストメッセージの形でクライアントから情報を受け取る
 2. HTTPリクエストメッセージを処理し、必要な作業を行う。
 3. HTMLを生成し、HTTPレスポンスメッセージに入れて返す。
結果として多くのWebアプリは全く性格の異なる「ハンドラ」と「テンプレートエンジン」という要素から構成される。

### ハンドラ
ハンドラは、クライアントから送られてきたHTTPリクエストを受け付け、処理する。  
また、HTMLを生成するためにテンプレートエンジンを呼び出し、最終的にはデータをHTMLレスポンスに組み込んで、クライアントに返送できるようにする。

### テンプレートエンジン
テンプレートとはHTMLに変換できるコードであり、クライアントにHTTPレスポンスメッセージを送り返すときに利用されるもの。  
テンプレートエンジンはテンプレートとデータから最終的なHTMLを生成する。

 * 静的テンプレート（ロジックレステンプレート）---HTMLにプレースホルダとなるトークンを埋め込んだもの。埋め込まれたトークンを正しいデータに置き換えることでHTMLを生成する
 * アクティブテンプレート---プレースホルダ・トークンだけでなく、条件式、繰り返し文、変数など、プログラミング言語の構造要素を含んでいる。

### アプリケーションのデザイン

#### Webアプリケーションの一般的な仕組み

1. クライアントがHTTPリクエストを送信
2. サーバがHTTPリクエストを処理
3. サーバがHTTPリクエストをクライアントに返信する

#### 典型的なWebアプリケーションでのサーバの働き
```
クライアント -> リクエスト -> マルチプレクサ -> ハンドラ
           <- レスポンス <-                    |
                                             V
                         テンプレート <-> テンプレートエンジン
```

### データモデル
今回のWebアプリは以下のデータ構造とし、リレーショナルデータベース（RDB）にマップされる  
* ユーザー---フォーラムのユーザー情報
* セッション---ユーザーが現在ログインしているセッション
* スレッド---フォーラムのスレッド（ユーザー間のやり取り）
* ポスト---スレッド内の投稿（ユーザーが追加したメッセージ）

### リクエストの受信と処理
1. クライアントはサーバのURLをターゲットとしてリクエストを送信する
2. サーバにはマルチプレクサがあり、処理を担当する所定のハンドラにリダイレクトする
3. ハンドラがリクエストを処理し、必要な仕事をする
4. ハンドラがテンプレートエンジンを呼び出し、クライアントに返信するための正しいHTMLを生成する

### クッキーを使ったアクセス制御
ユーザーがログインしたことがわかるように、レスポンスヘッダにクッキーを書く必要がある
クッキーはクライアントに送信され、ブラウザに保存される

