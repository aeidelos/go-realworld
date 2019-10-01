FROM golang:1.12.6 AS builder
WORKDIR /go/src/github.com/neotroops/go-realworld
COPY . .
RUN GO111MODULE=on go mod download
RUN GO111MODULE=on go mod verify
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/go-realworld main.go

FROM alpine:latest AS deployment
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/neotroops/go-realworld/build/go-realworld .
COPY --from=builder /go/src/github.com/neotroops/go-realworld/application.example.yaml .
CMD ["./go-realworld"]
