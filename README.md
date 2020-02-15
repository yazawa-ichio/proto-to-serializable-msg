# proto-to-serializable-msg
proto file compile simple message struct and MessagePack serializer

## 概要
protoファイルからメッセージ定義（クラス定義）とMessagePackの静的なシリアライザーを出力します。  
C#とJavaScript/TypeScript(NodeJs)とgolangで利用できます。
protocを利用しないためシングルバイナリで動作します。

### なにに使うの？
主にゲーム等のリアルタイム通信系でHTTPではなくUDP等を併用するケースで使用する目的で作成しました。  
後、protocの吐き出す定義ファイルが少し読みにくいので、シンプルにしたかったのも目的の一つです。  
（シリアライザー部分をつけたせいであんまり変わってないかも）  

### ProtocolBufferでよくない？
はい。ほぼ全てのケースで、そのままProtocolBufferを使う方がいいです。  
protoファイルを使うのであれば、gRPCを避ける必要がない気がするのでそのまま使いましょう。  

### How To Use

```
# get
go get github.com/yazawa-ichio/proto-to-serializable-msg/cmd/proto-to-serializable-msg
# gen message
proto-to-serializable-msg -lang cs -input ./input/proto -output ./out/proto
# use Config
proto-to-serializable-msg -c ./proto-config.yml
```

### Option

|Short|Long||
|---|---|---|
|-l|-lang|[generate language](#Lang)|
|-i|-input|input dir or *.proto file|
|-o|-output|output dir|
|-d|-dryrun|output dryrun|
|-c|-config|[generate use config.yml](#Config)|

#### Lang

|short|long|
|---|---|
|cs|csharp|
|js|javascript|
|ts|typescript|
|go|golang|

#### Config

[sample config](./tests/config/proto-config.yml)

```yml
input: "../proto"
lang:
  go:
    output: "./out/go/proto"
  csharp:
    output: "./out/cs/proto"
    property: false
    serializable: true
  js:
    output: "./out/ts/proto"
    use_ts: true
    disable_package_to_dir: false
```

## [CREDITS](./CREDITS)
Thanks to the developers of the packages used to create this software.

## [LICENSE](./LICENSE)
