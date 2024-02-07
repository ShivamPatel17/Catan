# Makefile for managing a Dockerized Go application using Docker Compose

# Docker Compose command
DC=docker-compose

# Start the application
up:
	$(DC) up --build -d

# Stop the application
down:
	$(DC) down

# Remove containers and images
clean:
	$(DC) down --rmi all

# Follow the logs
logs:
	$(DC) logs -f

# Utility command to access container shell
shell:
	docker exec -it catanapp /bin/sh

.PHONY: up down clean logs shell
