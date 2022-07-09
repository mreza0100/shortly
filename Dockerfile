FROM golang:1.18 as builder

# copy module, download and cache dependencies
WORKDIR /src/
COPY go.mod /src/go.mod
COPY go.sum /src/go.sum
RUN go mod download -x

# copy sources
ADD ./internal/ ./internal/
ADD ./pkg/ ./pkg/
ADD ./cmd/ ./cmd/
ADD ./.env ./.env


RUN mkdir ./build

# build
RUN  go build -o ./build/exec ./cmd

CMD ["./build/exec", "run"]
