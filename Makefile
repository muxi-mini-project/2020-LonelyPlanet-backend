all: gotool
	@go build -v .
clean:
	rm -f main
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
	gofmt -w .
help:
	@echo "make - compile the source code with local vendor"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt'"

.PHONY: clean gotool help
