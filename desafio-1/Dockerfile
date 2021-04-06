FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

COPY . /go/src

RUN cd /go/src

RUN go build -o desafio-1 main.go

EXPOSE 8000

CMD ["./desafio-1"]