run:
	go run cmd/api/main.go

compose:
	docker-compose -f ./build/docker-compose.yaml up -d

compose-rebuild:
	docker-compose -f ./build/docker-compose.yaml up -d --build --force-recreate