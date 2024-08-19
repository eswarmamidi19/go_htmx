run: build
	@./bin/go_htmx

build :
	@go build -o bin/go_htmx cmd/app/main.go