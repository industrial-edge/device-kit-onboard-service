FROM golang:1.20.6-bullseye

## Install unzip and jq
## Install lsb-release wget software-properties-common gnupg (These are releated with llvm) - for memory sanitizer tests
RUN apt-get update && apt-get install -y --no-install-recommends \
    unzip \
    jq \
    lsb-release \
    wget \
    software-properties-common \
    gnupg

# Install protobuf from releas page
RUN PB_REL="https://github.com/protocolbuffers/protobuf/releases" && VERSION=23.4 \
    && curl -LO $PB_REL/download/v$VERSION/protoc-$VERSION-linux-x86_64.zip \
    && unzip protoc-$VERSION-linux-x86_64.zip -d $HOME/.local \
    && export PATH="$PATH:$HOME/.local/bin" \
    && echo 'export PATH=$PATH:$HOME/.local/bin' >> ~/.bashrc

# Install LLVM as default compiler, for memory sanitizer
ENV LLVM_VERSION 15
RUN wget https://apt.llvm.org/llvm.sh && chmod +x llvm.sh && ./llvm.sh ${LLVM_VERSION}

## Set Clang as default CC

RUN echo "export CC=clang-${LLVM_VERSION}" >> ~/.bashrc && \
    . ~/.bashrc

# Sonar Scanner
ENV SONAR_SCANNER_VERSION 4.8.0.2856
RUN mkdir -p /progs && \
    # SonarScanner
    cd /progs && \
    wget https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-${SONAR_SCANNER_VERSION}-linux.zip && \
    unzip sonar-scanner-cli-${SONAR_SCANNER_VERSION}-linux.zip && \
    mv sonar-scanner-${SONAR_SCANNER_VERSION}-linux /opt/sonar-scanner && \
    echo 'export PATH=$PATH:/opt/sonar-scanner/bin' >> ~/.bashrc && \
    . ~/.bashrc

# JFrog Client
RUN curl -fL https://getcli.jfrog.io/v2 | sh && \
    mv jfrog /usr/bin/jfrog

# Install Deb/Rpm packaging tool 'goreleaser'
RUN echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | tee /etc/apt/sources.list.d/goreleaser.list
RUN apt update
RUN apt install goreleaser -y

# Setup Go environment
ENV GOPATH /go
ENV PATH ${GOPATH}/bin:$PATH

# Download go modules
ARG BASE_PACKAGE_NAME
WORKDIR /go/src/$BASE_PACKAGE_NAME/
COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on go mod download

# Install golint
RUN go get -d -v -u golang.org/x/lint/golint

# Install gRPC & Protobuf dependencies
RUN go get -d -v -u google.golang.org/grpc
RUN go get -d -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go get -d -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN go get -d -v -u github.com/golang/protobuf/protoc-gen-go

COPY . /go/src/$BASE_PACKAGE_NAME/
