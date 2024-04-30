.PHONY: build
build:
	mkdir -p build
	ogen ./swagger.yaml
	sqlc generate
	tinygo build -o ./build/app.wasm -target wasm -no-debug

.PHONY: build-go
build-go:
	mkdir -p build
	go build -o ./build/app.wasm

.PHONY: deploy
deploy:
	wrangler deploy

.PHONY: dev
dev:
	wrangler dev
