tests:
	go build -o bin/protoc-gen-msgpack
	protoc -I. --plugin=./bin/protoc-gen-msgpack --csharp_out=./test/output --js_out=import_style=commonjs,binary:./test/output  --msgpack_out=./test/msgpack ./test/proto/*.proto
