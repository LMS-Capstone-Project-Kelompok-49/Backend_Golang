FROM golang:1.17-buster
RUN apt-get update && \
    apt-get install -y git ca-certificates tzdata && \
    update-ca-certificates
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /dist

EXPOSE 3222

CMD ["/dist"]