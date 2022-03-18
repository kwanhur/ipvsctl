# Copyright 2022 kwanhur
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# init project path
WORKROOT := $(shell pwd)

# init environment variables
export PATH        := $(shell go env GOPATH)/bin:$(PATH)
export GO111MODULE := on

# init command params
GO           := go
GOBUILD      := $(GO) build
GOTEST       := $(GO) test
GOVET        := $(GO) vet
GOGET        := $(GO) get
GOGEN        := $(GO) generate
GOCLEAN      := $(GO) clean
GOINSTALL    := $(GO) install
GOFLAGS      := -race
STATICCHECK  := staticcheck
LICENSEEYE   := license-eye
MDLINT       := markdownlint
PIP          := pip3
PIPINSTALL   := $(PIP) install
NPM          := npm
NPMINSTALL   := $(NPM) install

# init ipvsctl version
VERSION ?= $(shell cat VERSION)
# init git commit id
COMMITID ?= $(shell git rev-parse --short HEAD)
# init built date
BUILT=`date`
# init built target
TARGET=ipvsctl
# init ldflags
LDFLAGS := "-s -w -X 'main.Version=v$(VERSION)' -X 'main.CommitID=$(COMMITID)' -X 'main.Built=$(BUILT)'"

# init arch
ARCH := $(shell getconf LONG_BIT)
ifeq ($(ARCH),64)
	GOTEST += $(GOFLAGS)
endif

# go install package
# $(1) package name
# $(2) package address
define INSTALL_PKG
	@echo installing $(1)
	$(GOINSTALL) $(2)
	@echo $(1) installed
endef

define PIP_INSTALL_PKG
	@echo installing $(1)
	$(PIPINSTALL) $(1)
	@echo $(1) installed
endef

define NPM_INSTALL_PKG
	@echo installing $(1)
	$(NPMINSTALL) $(1)
	@echo $(1) installed
endef

.PHONY: all
all:clean build

clean:
	rm -f $(TARGET)

build:
	$(GOBUILD) -ldflags $(LDFLAGS) -v -o $(TARGET)

# make deps, install dependent tools
deps:
	$(call NPM_INSTALL_PKG, markdownlint-cli)
	$(call PIPINSTALL_PKG, pre-commit)
	$(call INSTALL_PKG, license-eye, github.com/apache/skywalking-eyes/cmd/license-eye@latest)
	$(call INSTALL_PKG, staticcheck, honnef.co/go/tools/cmd/staticcheck@latest)

# make precommit, enable autoupdate and install hooks
precommit:
	pre-commit autoupdate
	pre-commit install
	pre-commit install-hooks

# make check
check:
	$(STATICCHECK) ./...

# make license-check, check code file's license declaration
license-check:
	$(LICENSEEYE) header check

# make license-fix, fix code file's license declaration
license-fix:
	$(LICENSEEYE) header fix

# make mdlint
mdlint:
	$(MDLINT) $(WORKROOT)

# make mdlint-fix
mdlint-fix:
	$(MDLINT) --fix $(WORKROOT)
