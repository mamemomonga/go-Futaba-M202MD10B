
doc.md: vfd.go vfd_test.go
# https://github.com/princjef/gomarkdoc
	gomarkdoc --output documents/godoc.md .

test:
	go test -v -count 1

.PHONY: test
