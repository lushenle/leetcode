ifneq (,)
.error This Makefile requires GNU Make.
endif

GO      = go
GOCOVER = $(GO) tool cover

test:
	$(GO) test -v -cover ./...

update:
	$(GO) get -u -v ./...

fmt:
	$(GO) fmt ./...

cover_test:
	$(GO) test -v ./... -covermode=count -coverprofile=coverage.out
	$(GOCOVER) -func=coverage.out -o=coverage.out

.PHONY: test update fmt cover_test
