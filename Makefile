BUILD=docker build --rm -q -f "Dockerfile" -t executable_tester:latest "."
RUN=docker run --rm=true -it --name testing --network scheduler-cluster --privileged executable_tester:latest
INSTALL=go install ./internal/executor
GOTEST=go test -count=1
RUNTESTALL=$(GOTEST) -tags debug ./...
RUNTESTSHORT=$(GOTEST) -tags debug,longTests ./...
RUNDEBUG=$(GOTEST) ./pkg/executable -run ^TestDebug

testAll:
	$(BUILD)
	$(RUN) make runInDockerAll

testShort:
	$(BUILD)
	$(RUN) make runInDockerShort

debug:
	$(BUILD)
	$(RUN) make debugInDocker

it:
	$(BUILD)
	$(RUN) /bin/bash


#DO NOT RUN THIS OUTSIDE OF DOCKER
runInDockerAll:
	$(INSTALL)
	$(RUNTESTALL)

#DO NOT RUN THIS OUTSIDE OF DOCKER
runInDockerShort:
	$(INSTALL)
	$(RUNTESTSHORT)

#DO NOT RUN THIS OUTSIDE OF DOCKER
debugInDocker:
	$(INSTALL)
	$(RUNDEBUG)

