FROM golang

ADD . /go/src/worldspan-simulator
#ADD src /go/src/worldspan-simulator
#ADD src/gorilla-mux /go/src/github.com/gorilla/mux

RUN go install worldspan-simulator

ENTRYPOINT /go/bin/worldspan-simulator
EXPOSE 8080
