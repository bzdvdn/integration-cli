package templates

// MakefileGO ...
func MakefileGO() (string, string) {
	filename := "Makefile"
	body := `.PHONY: build
build:
	go build -v 

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: run
run:
	go run main.go

.DEAFAULT_GOAL := run
`
	return filename, body
}
