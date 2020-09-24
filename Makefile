BUILD=docker build --rm -q -f "Dockerfile" -t executable_tester:latest "."
RUN=docker run --rm=true -it --name testing --network scheduler-cluster --privileged executable_tester:latest

test:
	$(BUILD)
	$(RUN) go test ./pkg/executable

it:
	$(BUILD)
	$(RUN) /bin/bash

