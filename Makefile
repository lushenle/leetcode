ifneq (,)
.error This Makefile requires GNU Make.
endif

test:
	go test -v -cover ./...

update:
	go get -u -v ./...

fmt:
	go fmt ./...

.PHONY: test update fmt
