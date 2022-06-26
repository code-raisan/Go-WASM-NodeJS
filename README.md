# Go-WASM NodeJS

NodeJSからGolangでWebAssemblyビルドした関数を呼び出すサンプルになります。

## About

Rustみたいに素直ではないGolangのWebAssemblyをなんとかしようとして、結局TSに移植する過程で負けたコードです。

デッドロックとかデッドロックとかデッドロックとかデッドロックとかに随分と時間を吸われてなんとか動くまでになったものです。

NodeJSから [jiro4989/ojosama](https://github.com/jiro4989/ojosama) を動作させるためにせっかくだしWebAssemblyで動かしたら速度改善も望めて
JSとの親和性も高くなるんじゃねと思い作っていたので中身の動作は、[jiro4989/ojosama](https://github.com/jiro4989/ojosama) となります

*`wasm_exec_node.js`と`wasm_exec.js`は公式が提供しているものを使用しています。*

*また`wasm_exec.js`はGoのWebAssemblyが特殊なためそのために必要となるものです。SDK的な何かだと思ってます(多分違うけど)*

## Run

> *Go 1.18.3*を使用してビルドしています

> *NodeJS v16.14.2*を使用して実行します

### 1. `main.go` をWASMでビルドします

```bash
# Windows (Power Shell 7.x)
$env:GOOS="js"
$env:GOARCH="wasm"
go build -o main.wasm .

# Linux
GOOS=js
GOARCH=wasm
go build -o main.wasm *.go
```

### 2. 実行

```bash
node ./wasm_exec_node.js ./main.wasm

# Result
[LOG] WebAssembly Initialized
[LOG] Converting
RESULT: おハーブですわ!
```

## License

ライセンスの所在はいまいちわからないので、すべて自分で書いた`main.go`はMITライセンスにしておきます。

[jiro4989/ojosama](https://github.com/jiro4989/ojosama) は読み込んでるだけだからここに書く必要があるのかはよくわからないが、[jiro4989/ojosama](https://github.com/jiro4989/ojosama)のライセンスに従います。

`wasm_exec_node.js`と`wasm_exec.js`は公式がBSDで提供してるんでまあって感じで

自由に使ってください(使い道あるか知らないけど、あったらスターください....)
