.PHONY: build test clean

GO=go

build:
	@$(GO) build -o ginie-darwin

test:
	@$(GO) test -v .

clean:
	@rm ginie-darwin
