FROM golang:1.20.1

## UPDATE THE OS
RUN apt-get update && \
    apt-get install protobuf-compiler -y && \
    apt-get install -y tzdata 

WORKDIR /go/src

## SET ENVIRONMENT
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV TZ America/Sao_Paulo

## COPY NECESSARY FILES
COPY go.* ./

## KEEP THE CONTAINER RUNNiNG
CMD ["tail", "-f", "/dev/null"]
