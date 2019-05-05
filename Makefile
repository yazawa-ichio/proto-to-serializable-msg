tests:
	go build -o bin/protoc-gen-msgpack
	mkdir -p ./test ./test/output ./test/msgpack
	protoc -I. --plugin=./bin/protoc-gen-msgpack \
		--csharp_out=./test/output \
		--js_out=import_style=commonjs_strict,binary:./test/output  \
		--msgpack_out=cs:./test/msgpack \
		--msgpack_out=js:./test/msgpack ./sample/proto/*.proto \

gen-sample:
	go build -o bin/protoc-gen-msgpack
	protoc -I./sample/ --plugin=./bin/protoc-gen-msgpack \
		--msgpack_out=js:./sample/server/proto \
		--msgpack_out=cs:./sample/client/proto ./sample/proto/*.proto \

run-sample: gen-sample;
	node ./sample/server/index.js

run-sample-client: gen-sample;
	cp -RT lib/cs/ ./sample/client/ProtoPack/
	dotnet run --project ./sample/client/client.csproj
