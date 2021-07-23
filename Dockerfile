FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go get -d -v .
RUN go install -v .

RUN go build -o ./server

EXPOSE 8000

CMD ["./server"]
