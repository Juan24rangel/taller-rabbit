FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY main.go .
RUN go build -o /servicio .

FROM alpine:3.20
COPY --from=build /servicio /servicio
CMD ["/servicio"]
