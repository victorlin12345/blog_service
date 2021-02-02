FROM golang:1.15 as builder
WORKDIR /builder/
ADD . /builder/
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine:3.12
COPY --from=builder /builder/app/ .
ENTRYPOINT ["/app"]

