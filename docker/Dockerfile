ARG DOCKER_REGISTRY=docker.io
ARG GO_VERSION=1.21

FROM ${DOCKER_REGISTRY}/golang:${GO_VERSION}-bullseye AS build


WORKDIR /app
ENV CGO_ENABLED=0\
    GOOS=linux\
    GOOARCH=amd64

RUN useradd -u 1001 -m nonroot

COPY ./go.mod ./go.sum ./


ARG GOPROXY=https://goproxy.cn
RUN go mod download

COPY ./internal ./internal
COPY ./main.go ./

RUN go build -o /app/bin


FROM build as test
RUN go test -v ./...

FROM scratch as run

COPY --from=build /app/bin /app/bin
COPY ./config.yaml /
COPY --from=build /etc/passwd /etc/passwd

USER 1001

EXPOSE 8080

ENTRYPOINT [ "/app/bin", "serve" ]
