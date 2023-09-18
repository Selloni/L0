.PHONY: run clean docker nats all

all: nats run

run:
	go run cmd/app/main.go

nats: clean
	go build cmd/nats/nats.go

docker:
	cd docker && docker compose up -d

clean:
	rm -rf main nats
	rm -rf .idea

