FROM golang:1.18.2-stretch

## UPDATE THE OS
RUN apt-get update && \
    apt-get install -y tzdata 

WORKDIR /go/src

## SET ENVIRONMENT
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV TZ America/Sao_Paulo

## START THE PROJECT
RUN go mod init github.com/renatospaka/grpc-calculator

## COPY NECESSARY FILES
COPY go.* ./

## INSTALL GRPC RELATED LIBRARIES 
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

## INSTALL MY STANDARD LIBRARIES, IF ANY
RUN go get github.com/stretchr/testify

## TIDY THE PROJECT
RUN go mod download && \
    go mod tidy

## KEEP THE CONTAINER RUNNiNG
CMD ["tail", "-f", "/dev/null"]