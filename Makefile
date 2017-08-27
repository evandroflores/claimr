
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
	go test -v ./...