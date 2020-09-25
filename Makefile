BUILD=docker build --rm -q -f "Dockerfile" -t executable_tester:latest "."
RUN=docker run --rm=true -it --name testing --network scheduler-cluster --privileged executable_tester:latest
INSTALL=go install ./internal/executor
RUNTEST=go test -tags debug ./...

test:
	$(BUILD)
	$(RUN) make runInDocker

it:
	$(BUILD)
	$(RUN) /bin/bash


#DO NOT RUN THIS OUTSIDE OF DOCKER
runInDocker:
	$(INSTALL)
	$(RUNTEST)

