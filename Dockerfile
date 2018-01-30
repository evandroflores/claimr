FROM golang:1.9.3
WORKDIR /go/src/github.com/evandroflores/claimr
COPY . .

RUN go build -o build/claimr main.go
RUN go install -v github.com/evandroflores/claimr
CMD ["claimr"]
