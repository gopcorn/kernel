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

# Rule 1: sdk imports nothing else from this repository
# Rule 2: kernel only imports sdk (and itself)
isolation:
	@bad=$$(cd sdk && go list -deps -f '{{if not .Standard}}{{.ImportPath}}{{end}}' ./... \
	  | grep '^$(MODPATH)/' | grep -Ev '^$(MODPATH)/sdk(/|$$)' || true); \
	[ -z "$$bad" ] || { echo "sdk imports other repo packages:"; echo "$$bad"; exit 1; }
	@bad=$$(cd kernel && go list -deps -f '{{if not .Standard}}{{.ImportPath}}{{end}}' ./... \
	  | grep '^$(MODPATH)/' | grep -Ev '^$(MODPATH)/(sdk|kernel)(/|$$)' || true); \
	[ -z "$$bad" ] || { echo "kernel imports forbidden repo packages:"; echo "$$bad"; exit 1; }
	@echo "isolation ok"
