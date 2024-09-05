FROM golang:1.23-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /build ./

FROM alpine:3.20 AS build-release-stage

WORKDIR /app

COPY --from=build-stage /build /build

ENTRYPOINT ["/build"]