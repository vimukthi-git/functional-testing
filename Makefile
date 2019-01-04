IMAGE_NAME?=centrifugeio/functional-testing

build-docker: ## Build Docker Image
build-docker:
	@echo "Building Docker Image"
	@docker build -t ${IMAGE_NAME}:latest .
	@echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
	@docker push ${IMAGE_NAME}:latest