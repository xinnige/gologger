
.PHONY: setup
PKG_NAME:=.
PKG_DIR:=$(PKG_NAME)

setup:
	GO111MODULE=on go get golang.org/x/tools/cmd/goimports
	GO111MODULE=on go get golang.org/x/tools/cmd/cover
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint
	GO111MODULE=on go get github.com/axw/gocov/gocov
	GO111MODULE=on go get github.com/AlekSi/gocov-xml
	GO111MODULE=on go get github.com/jstemmer/go-junit-report
	GO111MODULE=on go get github.com/golang/mock/gomock
	GO111MODULE=on go get github.com/stretchr/testify/assert
	GO111MODULE=on go get

.PHONY: fmt
fmt:
	set -x&&find . -name '*.go' -not -wholename './mock/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

.PHONY: lint
lint:
	GO111MODULE=on go get $(PKG_DIR)/...
	GO111MODULE=on golangci-lint run ./...

.PHONY: test
test:
	GO111MODULE=on go test $(PKG_DIR)/...

.PHONY: cover
cover:
	GO111MODULE=on go test -v -coverprofile=coverage.txt -covermode count $(PKG_DIR)/... | go-junit-report >report.xml
	GO111MODULE=on gocov test ./... | gocov-xml > coverage.xml
	find . -name coverage.xml|xargs -n1 sed -i -e 's#github.com/[^/]*/[^/]*/##g'
	find . -name coverage.xml|xargs -n1 sed -i -e 's#/go/src#./#g'

.PHONY: clean
clean:
	find . -name debug.test | xargs rm -f
	find . -name "coverage*" | xargs rm -f
	find . -name "*.xml" | xargs rm -f
	find . -name "*.log" | xargs rm -f
