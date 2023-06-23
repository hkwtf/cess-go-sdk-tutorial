test:
	go test -v ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	# rm everything except pdf in assets
	find ./assets/* -type f ! -name '*.pdf' -delete
