TARGET=ipvsctl
VERSION=1.0.0
BUILT=`date`
GOOS=linux
GOARCH=amd64
GO111MODULE=on

.PHONY: all
all:clean build

clean:
	rm -f $(TARGET)
build:
	$(eval COMMITID=$(shell git rev-parse --short HEAD))
	GOOS=$(GOOS) GOARCH=$(GOARCH) GO111MODULE=$(GO111MODULE) go build -ldflags "-s -w -X 'main.Version=v$(VERSION)' -X 'main.CommitID=$(COMMITID)' -X 'main.Built=$(BUILT)'" -v -o $(TARGET)