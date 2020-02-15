
tests: tests-config tests-serialize

tests-serialize: tests-gen
	cd tests/serialize/ts && npm install && tsc
	cd tests/serialize/go && GOOS=linux GOARCH=amd64 go build -o server
	cd tests/serialize && docker-compose build && docker-compose up -d
	cd tests/serialize/cs && dotnet test
	cd tests/serialize && docker-compose down

test-ts: tests-gen
	cd tests/serialize/ts && tsc

tests-gen: test-build
	rm -rf ./tests/serialize/cs/proto ./tests/serialize/go/proto ./tests/serialize/ts/proto
	mkdir -p ./tests/serialize/cs/proto ./tests/serialize/go/proto ./tests/serialize/ts/proto
	bin/proto-to-serializable-msg -l cs -i ./tests/proto -o ./tests/serialize/cs/proto
	bin/proto-to-serializable-msg -l go -i ./tests/proto -o ./tests/serialize/go/proto
	bin/proto-to-serializable-msg -l js -i ./tests/proto -o ./tests/serialize/ts/proto
	bin/proto-to-serializable-msg -l ts -i ./tests/proto -o ./tests/serialize/ts/proto

tests-config: test-build
	rm -rf ./tests/config/out
	cd ./tests/config && mkdir -p out/cs out/cs_prop out/go out/go_root out/js out/js_skip
	bin/proto-to-serializable-msg -c tests/config/proto-config.yml
	cd ./tests/config && go build go_build.go
	cd ./tests/config && npm install && ./node_modules/.bin/ts-node node.ts
	cp ./tests/config/cs_* ./tests/config/out/cs
	cp ./tests/config/cs_* ./tests/config/out/cs_prop
	cd ./tests/config/out/cs && dotnet build
	cd ./tests/config/out/cs_prop && dotnet build

gen-sample: test-build
	rm -rf ./sample/client/proto ./sample/server/proto ./sample/goclient/proto
	mkdir -p ./sample/client/proto ./sample/server/proto ./sample/goclient/proto
	bin/proto-to-serializable-msg -l cs -i ./sample/proto -o ./sample/client/proto
	bin/proto-to-serializable-msg -l js -i ./sample/proto -o ./sample/server/proto
	bin/proto-to-serializable-msg -l ts -i ./sample/proto -o ./sample/server/proto
	bin/proto-to-serializable-msg -l go -i ./sample/proto -o ./sample/goclient/proto

run-sample: gen-sample;
	node ./sample/server/index.js

run-sample-ts: gen-sample;
	tsc --project ./sample/server/
	node ./sample/server/dist/index.js

run-sample-client: gen-sample;
	dotnet run --project ./sample/client/client.csproj

run-sample-goclient: gen-sample;
	cd sample/goclient && go build && ./goclient

update-credits:
	gocredits -w .

test-build:
	cd cmd/proto-to-serializable-msg && go build -o ../../bin/proto-to-serializable-msg

