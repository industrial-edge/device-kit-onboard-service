FROM golang:1.17

ENV http_proxy http://194.138.0.5:9400
ENV HTTP_PROXY http://194.138.0.5:9400
ENV https_proxy http://194.138.0.5:9400
ENV HTTPS_PROXY http://194.138.0.5:9400

# Install protobuf from source
RUN apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool
RUN git clone https://github.com/google/protobuf.git && \
    cd protobuf && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install && \
    ldconfig && \
    make clean && \
    cd .. && \
    rm -r protobuf

# Install LLVM as default compiler, for memory sanitizer

## Add apt key for LLVM repository
RUN wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | apt-key add -

## Add LLVM apt repository
RUN echo "deb http://apt.llvm.org/stretch/ llvm-toolchain-stretch-5.0 main" | tee -a /etc/apt/sources.list

## Install clang from LLVM repository
RUN apt-get update && apt-get install -y --no-install-recommends \
    apt-utils \
    apt-transport-https \
    clang \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

## Set Clang as default CC
ENV set_clang /etc/profile.d/set-clang-cc.sh
RUN echo "export CC=clang-5.0" | tee -a ${set_clang} && chmod a+x ${set_clang}

# Sonar Scanner
RUN mkdir -p /progs && \
    # SonarScanner
    cd /progs && \
    wget https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.3.0.2102-linux.zip && \
    unzip sonar-scanner-cli-4.3.0.2102-linux.zip && \
    mv sonar-scanner-4.3.0.2102-linux /opt/sonar-scanner && \
    echo 'export PATH=$PATH:/opt/sonar-scanner/bin' >> ~/.bashrc && \
    . ~/.bashrc


# JFrog Client
RUN curl -L \
    --request GET  "https://jfrog.bintray.com/jfrog-cli-go/1.30.4/jfrog-cli-linux-amd64/jfrog" \
    --output /usr/bin/jfrog && \
    chmod u+x /usr/bin/jfrog

# Install Deb/Rpm packaging tool 'goreleaser'
RUN curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh -s -- -b /usr/local/bin/ "v0.131.1"


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
