PACKAGES = \
	github.com/mariusmagureanu/logo/tests \
	github.com/mariusmagureanu/logo 	

test:
	@for pkg in ${PACKAGES}; do \
		echo ; \
		echo "*** Test: $$pkg ***" ; \
		echo ; \
		go test -cover -coverprofile=coverage.out $$pkg || exit 1 ; \
		echo ; \
		go tool cover -func=coverage.out ; \
	done

