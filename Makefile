tests:
	go build -o bin/protoc-gen-msgpack
	mkdir -p ./test ./test/output ./test/msgpack
	protoc -I. --plugin=./bin/protoc-gen-msgpack \
		--csharp_out=./test/output \
		--js_out=import_style=commonjs,binary:./test/output  \
		--msgpack_out=./test/msgpack ./proto/*.proto
