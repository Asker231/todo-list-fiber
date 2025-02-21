dev:
	go run cmd/main.go
startdb:
	sudo docker compose up -d
stopdb:
	docker-compose down		
