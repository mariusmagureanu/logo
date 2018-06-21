PACKAGES = \
	github.com/mariusmagureanu/logo/tests \
	github.com/mariusmagureanu/logo 	

build:
	GOOS=linux GOARCH=amd64 go build  -ldflags "-s -w"

osxbuild:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w"

test:
	@for pkg in ${PACKAGES}; do \
		echo ; \
		echo "*** Test: $$pkg ***" ; \
		echo ; \
		go test -cover -coverprofile=coverage.out $$pkg || exit 1 ; \
		echo ; \
		go tool cover -func=coverage.out ; \
	done

