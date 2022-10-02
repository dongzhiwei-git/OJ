BRANCH := $(shell git branch --show-current)
COMMIT := $(shell git log -1 --format='%h')
DATE := $(shell date +%Y%m%d)
IMAGE_TAG := ${BRANCH}-${DATE}-${COMMIT}


build:
	go mod tidy \
	&& go build -o hgoj *.go

run:
	./hgoj

image-builder:
	docker build --file Dockerfile --tag hgoj:${IMAGE_TAG} .

run-image:
	docker run -p 8000:8000 hgoj:${IMAGE_TAG} 