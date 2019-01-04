IMAGE_NAME?=centrifugeio/functional-testing
TRAVIS_BRANCH?=`git rev-parse --abbrev-ref HEAD`
GIT_SHORT_COMMIT=`git rev-parse --short HEAD`
TIMESTAMP=`date -u +%Y%m%d%H`
TAG="${TRAVIS_BRANCH}-${TIMESTAMP}-${GIT_SHORT_COMMIT}"

build-docker: ## Build Docker Image
build-docker:
	@command -v dep >/dev/null 2>&1 || go get -u github.com/golang/dep/...
	@dep ensure
	@echo "Building Docker Image"
	@docker build -t ${IMAGE_NAME}:${TAG} .
	@docker tag "${IMAGE_NAME}:${TAG}" "${IMAGE_NAME}:latest"
	@echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
	@docker push ${IMAGE_NAME}:${TAG}
	@docker push ${IMAGE_NAME}:latest