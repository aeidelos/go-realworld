.PHONY: setup run test-api migrate

setup:
	go mod download

run:
	go run main.go

test-api:
	sh ./integration-test.sh