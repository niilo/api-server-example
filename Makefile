get-dev-dependencies:
	go get -u

vendor:
	docker run --rm -it -v $(PWD):/app -e "GOPATH=/app/vendor" -w /app golang:1 go get -d

create-swagger-document:
	swagger -apiPackage="github.com/niilo/api-server-example" -mainApiFile="github.com/niilo/api-server-example/main.go"

build-static-go-binary:
	docker run --rm -it -v $(PWD):/app -e "GOPATH=/app/vendor" -w /app golang:1 sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o app'

run:
	docker run --rm -it -v $(PWD):/app -p 8080:8080 iron/base ./app

build-docker-image:
	docker build -t niilo/api-server-example .

run-docker:
	docker run --rm -it -p 1323:1323 niilo/api-server-example

ebs-create-zip:
	zip api-server-example.zip app Dockerfile
