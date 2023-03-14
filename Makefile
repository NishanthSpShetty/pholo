.PHONY: proto

proto:
	@./compile.sh

run: proto
	@go run main.go

tidy:
	@go mod tidy

fmt:
	@gofumpt -w *.go

clean:
	@rm -rf ./pkg/generated
