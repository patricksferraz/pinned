FROM golang:1.17

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
    apt-get install build-essential protobuf-compiler librdkafka-dev -y && \
    go get google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go get google.golang.org/protobuf/cmd/protoc-gen-go && \
    go get github.com/spf13/cobra/cobra && \
    wget https://github.com/ktr0731/evans/releases/download/0.9.1/evans_linux_amd64.tar.gz && \
    tar -xzvf evans_linux_amd64.tar.gz && \
    mv evans ../bin && rm -f evans_linux_amd64.tar.gz \
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["tail", "-f", "/dev/null"]
