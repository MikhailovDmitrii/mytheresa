FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o promotions cmd/api/*

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/promotions /app/promotions

RUN mkdir -p /app/var && touch /app/var/products.db


RUN addgroup promotions
RUN adduser mytheresauser --ingroup promotions --disabled-password

RUN chown -R mytheresauser:promotions /app/var

USER mytheresauser:promotions

EXPOSE 8080

ENTRYPOINT ["/app/promotions"]