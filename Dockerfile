FROM golang:1.21.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /exoplanetservice

EXPOSE 8080

CMD ["/exoplanetservice"]
