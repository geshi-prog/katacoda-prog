## サードパーティ製のツールをインストールする
サーバーの負荷テストを行えるheyというツールがあるので実際にインストールして動かしてみましょう。

### go get
実際にインストールしてみましょう。
<br>`go get github.com/rakyll/hey`{{execute}}<br>
実行してみます。
<br> `hey https://www.golang.org`{{execute}}<br>

## プログラムをフォーマットする
実際に書いたコードをフォーマットすることができます。<br>
ビルドする前必須になります。<br>

### go fmt
ファイルのサイズにも注目してみてください。<br>
<br>`ll`{{execute}}<br>
`go fmt ./hello.go`{{execute}}<br>
`ll`{{execute}}<br>
少し軽くなったことが分かりますね！

## パッケージのドキュメント等を確認したい
goではfmtやnet等の様々なパッケージやメソッドを読み込みます。<br>
ドキュメントを読みたいときには表示してくれるコマンドがあります。

### go doc
<br>`go doc fmt.Println`{{execute}}

### このステップのポイント
* インストールしたいとき「go get」
* フォーマットしたいとき「go fmt」
* ドキュメントを読みたいとき「go doc」