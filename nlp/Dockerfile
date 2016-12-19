FROM golang:1.7.3
ADD . /go/src/github.com/docker/ddc-opencontrol/nlp
WORKDIR /go/src/github.com/docker/ddc-opencontrol/nlp
RUN go install
ENTRYPOINT ["/go/bin/nlp"]