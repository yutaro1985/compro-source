# テンプレートについての注意

## bufio.Scannerについて

`fmt.Scan`だと遅いのでこちらを使っている。
バッファサイズは最初が10^4、最大が10^6と決め打ちしている。
デフォルトだと2^16=65536なので、それ以上の長さの文字列を受け取ろうとするとTLEするため、バッファサイズの設定は必須。
バッファの最大値を大きくしすぎるとメモリ確保に時間をつかうせいで微妙に立ち上がりが遅くなるので適度な範囲で設定。
10^6を超えるような長さの文字列を受け取ろうとする場合うまく行かないかも？
その場合はreadlineを使うかバッファを大きくするか。

## ライブラリ挿入の仕方

[InsertTemplate](https://marketplace.visualstudio.com/items?itemName=yskoht.vscode-insert-template) を使う。

下記はドキュメントから抜粋。
変数が使えないため、
`リポジトリの配置ディレクトリ/Template`を`"insertTemplate.directory"`に設定する。

```:json
"insertTemplate.directory": "~/prg/competitive-programming",
"insertTemplate.ignore": ["**/*.md", "**/tmp/**"]
```
