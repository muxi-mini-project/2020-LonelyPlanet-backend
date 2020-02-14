all: gotool
	@go build -v .
clean:
	rm -f main
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
	gofmt -w .
ca:
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"
help:
	@echo "make - compile the source code with local vendor"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt'"
	@echo "make ca - generate ca files"

.PHONY: clean gotool ca help
