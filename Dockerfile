FROM golang:1.20-alpine as buildbase

RUN apk add build-base git

ARG CI_JOB_TOKEN

WORKDIR /go/src/gitlab.com/rarimo/rarimo-core

ENV GO111MODULE="on"
ENV CGO_ENABLED=1
ENV GOOS="linux"
ENV GOPRIVATE=gitlab.com/*
ENV GONOSUMDB=gitlab.com/*
ENV GONOPROXY=gitlab.com/*

RUN echo "machine gitlab.com login gitlab-ci-token password $CI_JOB_TOKEN" > ~/.netrc
# RUN  git config --global url."https://gitlab-ci-token:$CI_JOB_TOKEN@gitlab.com/".insteadOf https://gitlab.com/
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go mod vendor

RUN go build -o /usr/local/bin/rarimo-core gitlab.com/rarimo/rarimo-core/cmd/rarimo-cored



###

FROM alpine:3.9

RUN apk add --no-cache ca-certificates

COPY --from=buildbase /usr/local/bin/rarimo-core /usr/local/bin/rarimo-core

ENTRYPOINT ["rarimo-core"]
