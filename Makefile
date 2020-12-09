BUILDER=heroku/buildpacks:18
IMAGE_NAME=k8s-pipeline-authz
IMAGE_REPO=jasonblanchard/${IMAGE_NAME}
LOCAL_TAG=${IMAGE_REPO}
LATEST_TAG=${IMAGE_REPO}:latest
VERSION=$(shell git rev-parse HEAD)
VERSION_TAG=${IMAGE_REPO}:${VERSION}

build:
	pack build ${IMAGE_REPO} --builder ${BUILDER}
	docker tag ${IMAGE_REPO} ${VERSION_TAG}

push: build
	docker push ${LATEST_TAG}
	docker push ${VERSION_TAG}
