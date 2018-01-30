
check-env:
ifndef GOPATH
	@echo "Couldn't find the GOPATH env"
	@exit 1
endif

check-keys:
ifndef CLAIMR_DATABASE
	@echo "Couldn't find the CLAIMR_DATABASE env"
	@exit 1
endif

ifndef CLAIMR_TOKEN
	@echo "Couldn't find the CLAIMR_TOKEN env"
	@exit 1
endif

build: check-env
	@go build -o build/claimr main.go
	@echo "\nCheck the binary on the build dir build/claimr\n"
	@ls -lah build

install: check-env
	go install -v github.com/evandroflores/claimr

run: check-env check-keys
	@go run main.go

docker-build:
	docker build -t evandroflores/claimr .

docker-run: check-keys
	@docker run -e "CLAIMR_TOKEN=${CLAIMR_TOKEN} CLAIMR_DATABASE=${CLAIMR_DATABASE}" evandroflores/claimr

test: check-keys
	@go test -gcflags=-l -cover ./...

cover: check-keys
	@rm -f coverage.*
	@echo 'mode: atomic' > coverage.txt
	go list ./... | grep -v vendor | xargs -n1 -I{} sh -c 'touch coverage.out & go test -gcflags=-l -race -covermode=atomic -coverprofile=coverage.out {} && tail -n +2 coverage.out >> coverage.txt'

open-cover:
	go tool cover -html coverage.txt
