clean: 
	@go clean
	@rm -rf ./bin

build: clean	
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/task app/task/src/main.go

start: build
	sls offline --useDocker --noPrependStageInUrl