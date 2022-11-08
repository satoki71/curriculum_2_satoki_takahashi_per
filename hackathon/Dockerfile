# syntax=docker/dockerfile:1

FROM golang:1.19 as build

WORKDIR /go/src

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./*.go  ./

ARG CGO_ENABLED=0
#ARG GOOS=linux
#ARG GOARCH=amd64
RUN go build \
    -o /go/bin/docker-hackathon \
    -ldflags '-s -w'

FROM gcr.io/distroless/base-debian10 as runner

WORKDIR /

COPY --from=build /go/bin/docker-hackathon /app/docker-hackathon

EXPOSE 8000

#USER nonroot:nonroot

ENTRYPOINT ["/app/docker-hackathon"]

#to run, "docker run --publish 8000:8000 docker-cd:multistage"