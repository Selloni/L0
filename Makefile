.PHONY: run clean docker

run:
	go build cmd/app/main.go
	./main

nats: clean
	go build cmd/nats/nats.go
docker:
	cd docker && docker compose up

clean:
	rm -rf main nats
	rm -rf .idea

