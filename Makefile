SERVICE_NAME=noname-grpc

.PHONY: build
build:
	@running=$$(docker-compose ps $(SERVICE_NAME) | grep -c "Up"); \
	if [ "$$running" -eq 0 ]; then \
		echo "Container not running, restarting..."; \
		docker-compose up; \
	fi;

shell:
	@docker-compose exec $(SERVICE_NAME) bash
