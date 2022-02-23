include common/make/common.makefile

# Ensure the service name matches the name of the repo
SERVICE=my-new-service

# These are using for semantic versioning for Go binaries, and dockerfiles
# We use BUILD_NUMBER in place of PATCH for tracibility between CI tooling, and versioning
# This means that build number will increase indefinitely, alongside the job number in the CI tool
MAJOR=0
MINOR=0
BUILD_NUMBER=CI-TOOL-SETS-THIS
SEMVER?="$(MAJOR).$(MINOR).$(BUILD_NUMBER)-$(GIT_COMMIT)"

GIT_COMMIT=$(shell git log -1 --pretty=%h)
GIT_COMMIT_PREV=$(shell git log -2 --pretty=%h)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GO_VERSION=$(shell go version)
GO_TAG="v$(MAJOR).$(MINOR).$(BUILD_NUMBER)"
DOCKER_IMAGE_NAME="$(SERVICE)-$(BRANCH):$(SEMVER)"

# DOCKER_REGISTRY if set, must end in /
DOCKER_REGISTRY?= 



