FROM golang:1.23-alpine as buildbase

RUN apk add build-base git

WORKDIR /go/src/github.com/rarimo/rarimo-core

ENV GO111MODULE="on"
ENV CGO_ENABLED=1
ENV GOOS="linux"

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go mod vendor
# TODO switch to latest cosmosvisor
RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.5.0

RUN cp $GOPATH/bin/cosmovisor /usr/local/bin/cosmovisor
RUN go build -mod=mod -o /usr/local/bin/rarimo-core github.com/rarimo/rarimo-core/cmd/rarimo-cored

FROM alpine:3.9

RUN apk add --no-cache ca-certificates

COPY --from=buildbase /usr/local/bin/rarimo-core /usr/local/bin/rarimo-core
COPY --from=buildbase /usr/local/bin/cosmovisor /usr/local/bin/cosmovisor

ENTRYPOINT ["rarimo-core"]