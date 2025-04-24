FROM golang:1.24-alpine

WORKDIR /app

ARG HOST
ARG PORT

ENV APP_HOST=$HOST
ENV APP_PORT=$PORT

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o myapp ./cmd/api/main.go

EXPOSE ${PORT}

CMD ["./myapp"]
