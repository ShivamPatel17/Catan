# Docker Compose command
DC=docker-compose

# Image names and tags
APP_IMAGE_NAME=catan-app
APP_IMAGE_TAG=latest
APP_IMAGE_FILE=$(APP_IMAGE_NAME)-$(APP_IMAGE_TAG).tar
WEB_IMAGE_NAME=catan-web
WEB_IMAGE_TAG=latest
WEB_IMAGE_FILE=$(WEB_IMAGE_NAME)-$(WEB_IMAGE_TAG).tar

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

# Save images - run this to save your latest build "offline"
save-images:
	docker save $(APP_IMAGE_NAME):$(APP_IMAGE_TAG) > $(APP_IMAGE_FILE)
	docker save $(WEB_IMAGE_NAME):$(WEB_IMAGE_TAG) > $(WEB_IMAGE_FILE)

# Load images - run this when you're offline and want to load up the latest images load-images:
	docker load < $(APP_IMAGE_FILE)
	docker load < $(WEB_IMAGE_FILE)

# Run the services in offline mode - make sure to load up images
offline:
	$(DC) -f docker-compose.yml -f docker-compose.offline.yml up --pull never -d
