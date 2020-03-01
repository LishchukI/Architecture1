default: out/example

clean:
	rm -rf out

test:
	go vet && go test

out/example: implementation.go cmd/example/main.go
	mkdir -p out
	echo package main > cmd/example/version.go
	echo -n "var BuildVersion = \"`git describe`\"" >> cmd/example/version.go
	go build -o out/example ./cmd/example