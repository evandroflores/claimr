
check-env:
ifndef GOPATH
	@echo "Couldn't find the GOPATH env"
	@exit 1
endif

check-token:
ifndef CLAIMR_TOKEN
	@echo "Couldn't find the CLAIMR_TOKEN env"
	@exit 1
endif

build: check-env vendorize
	@go build -o build/claimr main.go
	@echo "\nCheck the binary on the build dir build/claimr\n"
	@ls -lah build


run: check-env check-token
	@go run main.go

docker-build:
	docker build -t evandroflores/claimr .

docker-run: check-token
	docker run -e CLAIMR_TOKEN=${CLAIMR_TOKEN} evandroflores/claimr

test:
	go test -cover ./...

cover:
	echo 'mode: atomic' > coverage.txt && go list ./... | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp
	go tool cover -html=coverage.txt -o coverage.html
