.PHONY: run clean docker

run:
	go build cmd/app/main.go
	./main

nats:
	go run cmd/nats/nats.go

docker:
	cd docker && docker compose up

clean:
	rm -rf main
	rm -rf .idea

