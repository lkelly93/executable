BUILD=docker build --rm -q -f "Dockerfile" -t executable_tester:latest "."
RUN=docker run --rm=true -it --name testing --network scheduler-cluster --privileged executable_tester:latest
INSTALL=go install ./internal/executor
RUNTESTALL=go test -tags debug ./...
RUNTESTSHORT=go test -tags debug,longTests ./...
RUNDEBUG=go test ./pkg/executable -run ^TestDebug

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

