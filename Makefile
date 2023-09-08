.PHONY: run clean

run:
	go build cmd/app/main.go
	./main

clean:
	rm -rf "main
	rm -rf .idea
