FROM golang

ADD . /go/src/worldspan-simulator
ADD data /go/data
ADD htmlTemplates /go/htmlTemplates

RUN go get worldspan-simulator
RUN go install worldspan-simulator

ENTRYPOINT /go/bin/worldspan-simulator
EXPOSE 8080
