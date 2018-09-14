deploy: build/client build/go
	now --public --docker && now alias

build/client:
	yarn build:prod

build/go:
	env GOOS=linux env GOARCH=amd64 env CGO_ENABLED=0 go build -o ./cmd/main main.go