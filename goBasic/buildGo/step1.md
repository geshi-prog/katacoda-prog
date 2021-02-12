## HelloWorldを実行する
まずは、インストールされているGoのバージョンを確認してみます。
`go version`{{execute}}<br>
次にhello.goを作成してみましょう。GOのプログラムファイルの拡張子は.goになります。`hello.go`{{open}}<br>
それでは、実際に動かしてみます。

### go run
<br>`go run ./hello.go`{{execute}}<br>
実際に動くことを確認出来たらバイナリファイルにビルドしてみます。

### go build
<br>`go build ./hello.go`{{execute}}<br>
実際に作成されたかを確認してみます。
<br>`ls`{{execute}}<br>
helloが出来てれば大丈夫です！<br>
それでは、作成されたhelloが動くか確認します。
<br>`./hello`{{execute}}<br>
作成するファイル名を指定したい場合「-o」を使います。
<br>`go build -o hello_world hello.go`{{execute}}<br>
`ls`{{execute}}<br>
`./hello_world`{{execute}}<br>

### このステップのポイント
* goの拡張子は「.go」
* プログラムを実行する時は、「go run ○○.go」
* プログラムをビルドする時は、「go build ○○.go」

次のステップではサードパーティのツールのインストールなどを行います。