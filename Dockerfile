# syntax=docker/dockerfile:1
FROM golang:1.21.5-alpine AS build

WORKDIR /usr/src/app

#install go-swagger
RUN go install github.com/go-swagger/go-swagger/cmd/swagger@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -C cmd/api-server -v -o /usr/local/bin/madr

# generate swagger file
RUN swagger generate spec -o ./swagger.yaml --scan-models

FROM gcr.io/distroless/static-debian12

WORKDIR /home/nonroot/

USER nonroot:nonroot

COPY --from=build /usr/src/app/swagger.yaml /usr/local/bin/madr ./

ENTRYPOINT ["./madr"]