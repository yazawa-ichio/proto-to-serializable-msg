# (WIP)protoc-gen-msgpack
proto file compile simple message struct and MessagePack serializer

## 作業中
protoファイルからメッセージ定義（クラス定義）とMessagePackの静的なシリアライザーを出力します。  
現在、C#とNodeJs用を概ね作成。あとはTypeScriptとgolang用を作成予定。  
golangとNodeJsとTypeScriptの勉強目的で作成中。  

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
go get github.com/yazawa-ichio/protoc-gen-msgpack
# gen message
protoc -I. --msgpack_out=cs:./output/cs --msgpack_out=js:./output/js *.proto
```


