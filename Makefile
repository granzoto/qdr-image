PROJECT_NAME=qdrouterd
DOCKER_REGISTRY=quay.io
DOCKER_ORG=interconnectedcloud
PWD=$(shell pwd)

 # This is the latest version of the Qpid Dispatch Router
DISPATCH_VERSION=1.12.0
DISPATCH_SNAPSHOT_VERSION=1.13.0
PROTON_VERSION=0.31.0
PROTON_SOURCE_URL=http://archive.apache.org/dist/qpid/proton/${PROTON_VERSION}/qpid-proton-${PROTON_VERSION}.tar.gz
ROUTER_SOURCE_URL=http://archive.apache.org/dist/qpid/dispatch/${DISPATCH_VERSION}/qpid-dispatch-${DISPATCH_VERSION}.tar.gz

# If a DOCKER_TAG is specified, go ahead and use it.
# if DOCKER_TAG is not specified use the DISPATCH_VERSION as the DOCKER_TAG
ifneq ($(strip $(DOCKER_TAG)),)
	DOCKER_TAG_VAL=$(DOCKER_TAG)
else
	DOCKER_TAG_VAL=$(DISPATCH_VERSION)
endif

all: build

build:
	docker build -t qdrouterd-builder:${DOCKER_TAG_VAL} builder
	docker run -ti -v $(PWD):/build:z -w /build qdrouterd-builder:${DOCKER_TAG_VAL} bash build_tarballs ${ROUTER_SOURCE_URL} ${PROTON_SOURCE_URL}

build-snapshot:
	docker build -t qdrouterd-builder:${DISPATCH_SNAPSHOT_VERSION}-snapshot builder
	docker run -ti -v $(PWD):/build:z -w /build qdrouterd-builder:${DISPATCH_SNAPSHOT_VERSION}-snapshot bash build_tarballs_snapshot master master

clean:
	rm -rf proton_build proton_install qpid-dispatch.tar.gz qpid-dispatch-src qpid-proton.tar.gz qpid-proton-src staging build

cleanimage:
	docker image rm -f qdrouterd-builder

buildimage:
	docker build -t ${PROJECT_NAME}:latest .
	docker tag ${PROJECT_NAME}:latest ${DOCKER_REGISTRY}/${DOCKER_ORG}/${PROJECT_NAME}:${DOCKER_TAG_VAL}

buildimage-snapshot: build-snapshot
	docker build -t ${PROJECT_NAME}:snapshot .
	docker tag ${PROJECT_NAME}:snapshot ${DOCKER_REGISTRY}/${DOCKER_ORG}/${PROJECT_NAME}:${DISPATCH_SNAPSHOT_VERSION}-snapshot

push-common:
# DOCKER_USER and DOCKER_PASSWORD is useful in the CI environment.
# Use the DOCKER_USER and DOCKER_PASSWORD if available
# if not available, assume the user has already logged in
ifneq ($(strip $(DOCKER_USER)$(DOCKER_PASSWORD)),)
	@docker login -u ${DOCKER_USER} -p ${DOCKER_PASSWORD} ${DOCKER_REGISTRY}
endif

push: buildimage push-common
	docker push ${DOCKER_REGISTRY}/${DOCKER_ORG}/${PROJECT_NAME}:${DOCKER_TAG_VAL}

push-snapshot: buildimage-snapshot push-common
	docker push ${DOCKER_REGISTRY}/${DOCKER_ORG}/${PROJECT_NAME}:${DISPATCH_SNAPSHOT_VERSION}-snapshot

.PHONY: build buildimage cleanimage clean push
