FROM golang:1.18-alpine as buildbase

WORKDIR /go/src/gitlab.com/rarify-protocol/rarimo-core
COPY vendor .
COPY . .

ENV GO111MODULE="on"
ENV CGO_ENABLED=1
ENV GOOS="linux"

RUN apk add build-base
RUN go build -o /usr/local/bin/rarimo-core gitlab.com/rarify-protocol/rarimo-core/cmd/rarimo-cored

###

FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/rarimo-core /usr/local/bin/rarimo-core
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["rarimo-core"]
