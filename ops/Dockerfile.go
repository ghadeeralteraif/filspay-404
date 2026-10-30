FROM golang:1.23-alpine AS build
WORKDIR /src
COPY apps/settlement/ .
RUN go build -o /settlement ./cmd/settlement

FROM alpine:3.20
WORKDIR /app
COPY --from=build /settlement ./settlement