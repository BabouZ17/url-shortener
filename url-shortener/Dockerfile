FROM golang:1.18.4-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /url-shortener

CMD [ "/url-shortener" ]