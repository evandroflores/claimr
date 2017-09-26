
check-env:
ifndef GOPATH
	@echo "Couldn't find the GOPATH env"
	@exit 1
endif

check-keys:
ifndef CLAIMR_TOKEN
	@echo "Couldn't find the CLAIMR_TOKEN env"
	@exit 1
endif

ifndef AWS_ACCESS_KEY_ID
	@echo "Couldn't find the AWS_ACCESS_KEY_ID env"
	@exit 1
endif

ifndef AWS_SECRET_ACCESS_KEY
	@echo "Couldn't find the AWS_SECRET_ACCESS_KEY env"
	@exit 1
endif

build: check-env vendorize
	@go build -o build/claimr main.go
	@echo "\nCheck the binary on the build dir build/claimr\n"
	@ls -lah build


run: check-env check-keys
	@go run main.go

docker-build:
	docker build -t evandroflores/claimr .

docker-run: check-keys
	@docker run -e CLAIMR_TOKEN=${CLAIMR_TOKEN} \
	            -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
	            -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
	            evandroflores/claimr

test:
	go test -cover ./...

cover:
	echo 'mode: atomic' > coverage.txt && go list ./... | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp
	go tool cover -html=coverage.txt -o coverage.html
