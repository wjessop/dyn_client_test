all:
	$(GOPATH)/bin/goop go build -o bin/test_client

clean:
	rm -f bin/*
	rm -rf .vendor

deps:
	$(GOPATH)/bin/goop install

update_deps:
	$(GOPATH)/bin/goop update
