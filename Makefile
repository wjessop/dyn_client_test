all:
	$(GOPATH)/bin/goop go build -o bin/test_client

clean:
	rm -f bin/*

deps:
	$(GOPATH)/bin/goop install

update_deps:
	$(GOPATH)/bin/goop update
