.PHONY: setup run test-api migrate

setup:
	export GO111MODULE=on 
	go mod download

run:
	go run main.go

test-api:
	sh ./integration-test.sh