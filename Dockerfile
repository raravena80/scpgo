FROM alpine:latest
MAINTAINER Ricardo Aravena <raravena@branch.io>

ENV PATH $PATH:/go/bin:/usr/local/go/bin
ENV GOPATH /go

RUN	apk add --no-cache ca-certificates

COPY . /go/src/github.com/raravena80/scpgo

RUN set -x \
	&& apk add --no-cache --virtual .deps \
		gcc libc-dev git libgcc go \
	&& cd /go/src/github.com/raravena80/scpgo \
        && go get ./... \
	&& CGO_ENABLED=0 go build -o /usr/bin/scpgo . \
	&& apk del .deps \
	&& rm -rf /go \
	&& echo "Build Finished."

ENTRYPOINT [ "scpgo" ]
