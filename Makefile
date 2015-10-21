get-dev-dependencies:
	go get -u

vendor:
	docker run --rm -it -v $(PWD):/app -e "GOPATH=/app/vendor" -w /app golang:1 go get -d

generate-swagger:
	mkdir -p sw-generate
	cp *.go sw-generate
	swagger -apiPackage="github.com/niilo/api-server-example/sw-generate" -mainApiFile="github.com/niilo/api-server-example/sw-generate/server.go"
	rm -rf sw-generate

build-static-binary:
	docker run --rm -it -v $(PWD):/app -e "GOPATH=/app/vendor" -w /app golang:1 sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o app'

run-in-docker:
	docker run --rm -it -v $(PWD):/app -w /app -p 1323:1323 golang:1 ./app

docker-image-build:
	docker build -t niilo/api-server-example .

docker-image-run:
	docker run --rm -it -p 1323:1323 niilo/api-server-example

ebs-create-zip:
	zip api-server-example.zip app Dockerfile
