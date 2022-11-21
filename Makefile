build:
	@go build -o ./bin/reverse-proxy

build_mac_amd:
	@env GOOS=darwin GOARCH=amd64 go build -o bin/mac/amd/pubctl

build_mac_arm:
	@env GOOS=darwin GOARCH=arm64 go build -o bin/mac/arm/pubctl

build_win_amd64:
	@env GOOS=windows GOARCH=amd64 go build -o bin/win/amd64/pubctl.exe

build_win_386:
	@env GOOS=windows GOARCH=386 go build -o bin/win/386/pubctl.exe

build_linux_amd64:
	@env GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/pubctl

run: build
	@./bin/reverse-proxy

test:
	@go test -v ./...