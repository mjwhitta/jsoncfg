all: fmt

check:
	@which go >/dev/null 2>&1

clean: fmt

clena: clean

fmt: check
	@go fmt .
