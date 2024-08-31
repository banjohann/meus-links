run:
	@docker compose up -d
	@go run cmd/meus-links/main.go