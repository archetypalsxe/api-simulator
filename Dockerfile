FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/api-simulator/app
COPY data /go/src/api-simulator/data
COPY htmlTemplates /go/src/api-simulator/htmlTemplates
COPY css /go/src/api-simulator/css
COPY js /go/src/api-simulator/js
WORKDIR /go/src/api-simulator/app

RUN go get ./
RUN go build -o api-simulator

CMD if [ ${APP_ENV} = production ]; \
    then \
        api-simulator; \
    else \
        go get github.com/pilu/fresh && \
        fresh; \
    fi

EXPOSE 8080
