FROM golang

ADD . /go/src/github.com/golang/ghanto/sds011-server

RUN go install -v github.com/golang/ghanto/sds011-server

ENTRYPOINT /go/bin/sds011-server

EXPOSE 9099