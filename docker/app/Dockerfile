# Builder container
ARG GO_VERSION
ARG MODULE

FROM golang:${GO_VERSION}-alpine as build

ENV GO111MODULE=on

WORKDIR /go/src/${MODULE}
ADD ./ /go/src/${MODULE}

RUN apk add --no-cache git
RUN go build -o app
RUN go get gopkg.in/urfave/cli.v2@master
RUN go get github.com/oxequa/realize
RUN go get -u github.com/go-delve/delve/cmd/dlv
RUN go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

# Application container
FROM alpine:latest

WORKDIR /approot
COPY --from=build /go/src/${MODULE}/app .

RUN addgroup go \
RUN adduser -D -G go go \
RUN chown -R go:go /approot/app

CMD ["./app"]
