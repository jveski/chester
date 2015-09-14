test:
	go test ./...

release: test
	go get github.com/mitchellh/gox
	gox -build-toolchain
	gox -output "dist/{{.OS}}_{{.Arch}}"

clean:
	rm -rf dist
