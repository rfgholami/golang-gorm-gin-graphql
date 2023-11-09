FROM golang:1.21.3

WORKDIR /go/src/github.com/kwa0x2/GoLang-Postgre-API

COPY . /go/src/github.com/kwa0x2/GoLang-Postgre-API/

RUN go get -v

RUN go build -o main .

CMD [ "./main" ]