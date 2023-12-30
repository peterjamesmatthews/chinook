.PHONY: dev
dev: 
	docker compose -f docker-compose.yml up -d --build
	docker image prune -f

.PHONY: clean
clean: 
	docker compose -f docker-compose.yml down -t 1
