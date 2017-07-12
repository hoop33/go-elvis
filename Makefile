PACKAGES ?= .

default: build

.PHONY: build
build: check
	go build

.PHONY: check 
check: vet lint errcheck interfacer structcheck varcheck unconvert gosimple staticcheck unused test

.PHONY: vet 
vet:
	go vet $(PACKAGES)

.PHONY: lint 
lint:
	golint -set_exit_status $(PACKAGES)

.PHONY: errcheck 
errcheck:
	errcheck $(PACKAGES)

.PHONY: interfacer 
interfacer:
	interfacer $(PACKAGES)

.PHONY: aligncheck 
aligncheck:
	aligncheck $(PACKAGES)

.PHONY: structcheck 
structcheck:
	structcheck $(PACKAGES)

.PHONY: varcheck
varcheck:
	varcheck $(PACKAGES)

.PHONY: unconvert 
unconvert:
	unconvert -v $(PACKAGES)

.PHONY: gosimple
gosimple:
	gosimple $(PACKAGES)

.PHONY: staticcheck 
staticcheck:
	staticcheck $(PACKAGES)

.PHONY: unused
unused:
	unused $(PACKAGES)

.PHONY: test
test:
	go test -cover $(PACKAGES)

.PHONY: clean
clean:
	go clean

.PHONY: deps
deps:
	go get -u github.com/FiloSottile/vendorcheck
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
	go get -u github.com/mdempsky/unconvert
	go get -u github.com/mvdan/interfacer/cmd/interfacer
	go get -u github.com/opennota/check/...
	go get -u github.com/yosssi/goat/...
	go get -u honnef.co/go/tools/...
