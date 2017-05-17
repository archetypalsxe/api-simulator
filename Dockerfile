FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/api-simulator/app
COPY data /go/data
COPY htmlTemplates /go/htmlTemplates
WORKDIR /go/src/api-simulator/app

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
    then \
        app; \
    else \
        go get github.com/archetypalsxe/api-simulator && \
        api-simulator; \
    fi

EXPOSE 8080
