run : build
	@./bin/cmd/api/api
build :
	@go build -o ./bin/cmd/api/api ./cmd/api/api.go