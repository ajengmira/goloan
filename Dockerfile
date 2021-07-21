FROM golang:1.12-alpine

WORKDIR /go/src/goloan

COPY *.go .

RUN go build -o main .

EXPOSE 8888

CMD [ "/main" ]