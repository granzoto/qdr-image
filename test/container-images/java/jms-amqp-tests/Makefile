TAG=1.0
IMAGE=docker.io/fgiorgetti/jms-amqp-tests

all: build-image push-image

build-image:
	docker build -t $(IMAGE):$(TAG) .

push-image:
	docker push $(IMAGE):$(TAG)
