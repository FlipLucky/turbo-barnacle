build:
	go build -o bin/app app/main.go

run: 
	go run bin/app

dev:	
	docker compose up -d && docker compose logs -f app
