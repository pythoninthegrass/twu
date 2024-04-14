# syntax=docker/dockerfile:1.6

FROM golang:1.22.2-alpine AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=arm64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /twu

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=builder /twu .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/twu"]
