FROM golang:latest

RUN apt-get update && \
    apt-get -y install vim
RUN mkdir /go/src/mongo-graphql
WORKDIR /go/src/mongo-graphql
COPY . /go/src/mongo-graphql

RUN go get github.com/99designs/gqlgen
RUN go get github.com/globalsign/mgo
#CMD ["./build/build"]