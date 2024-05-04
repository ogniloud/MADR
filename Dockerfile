# syntax=docker/dockerfile:1
FROM golang:1.22.0-alpine AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go test -v ./...
RUN go build -C cmd/api-server -v -o /usr/local/bin/madr

FROM gcr.io/distroless/static-debian12

WORKDIR /home/nonroot/

USER nonroot:nonroot

COPY --from=build /usr/src/app/swagger.yaml /usr/local/bin/madr ./

ENTRYPOINT ["./madr"]