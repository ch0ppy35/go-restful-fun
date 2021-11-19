FROM golang:1.17-alpine as builder

WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o ./go-resftul

FROM alpine:3.14 as package
WORKDIR /app
COPY --from=builder /build/go-resftul .
ENTRYPOINT [ "./go-resftul" ]