# docker build -t ${PWD##*/} .
# 1. compile GO code inside 'builder' container
FROM golang:1.9.3-alpine3.7 as builder
ENV APPNAME=go-examples GOPATH=/tmp

COPY docker/Proxy.traiana.int /usr/local/share/ca-certificates/proxy.pem
RUN  ls -al /usr/local/share/ca-certificates && \
    update-ca-certificates && apk --no-cache add ca-certificates git

COPY src/${APPNAME} /tmp/src/${APPNAME}
COPY docker/dep-linux-amd64 /tmp/bin/dep

WORKDIR /tmp/src/${APPNAME}
RUN ls -al /tmp/src/${APPNAME}/*
RUN chmod +x /tmp/bin/* && pwd && /tmp/bin/dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /tmp/app cmd/main.go

# 2. copy compiled app to final container
FROM alpine:3.7
COPY --from=builder /tmp/app /opt/app
COPY docker/tini-static /opt/tini
RUN  chmod +x /opt/tini && apk --no-cache add ca-certificates
WORKDIR /opt/
CMD ["/opt/tini","--","/opt/app"]