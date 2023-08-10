FROM golang:latest as build
WORKDIR /build
COPY . /build
RUN go build ./cmd/main.go

FROM alpine:latest as prod
COPY --from=build /build .
CMD ["./main"]