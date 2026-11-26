MODULES := $(shell go list -m -f '{{.Dir}}')

.PHONY: init vet build test isolation

init:
	@rm -f go.work go.work.sum
	go work init 
	go work use -r .
	@echo "go.work created:" && go list -m

vet:
	@for m in $(MODULES); do (cd $$m && go vet ./...) || exit 1; done

build:
	GOWORK=off go build ./...

test:
	@for m in $(MODULES); do echo "== $$m"; (cd $$m && go test -race ./...) || exit 1; done

# sdk imports nothing from this repository; kernel imports only sdk.
isolation:
	@for m in kernel; do \
	  deps=$$(cd $$m && go list -deps ./...) || { echo "go list failed in $$m"; exit 1; }; \
	  bad=$$(echo "$$deps" | grep -E 'gopcorn/(client|adapters|conformance|testfeatures)' || true); \
	  [ -z "$$bad" ] || { echo "kernel imports forbidden packages:"; echo "$$bad"; exit 1; }; \
	done; echo "isolation ok"
