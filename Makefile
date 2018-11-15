.PHONY: all
all:
	@echo "Skipping..."

.PHONY: get_dep
get_dep:
	@echo ">> getting dependencies"
	go get -t ./...

.PHONY: test
test: get_dep
	go test

.PHONY: lint
lint:
	@gofmt -s -w coindesk
