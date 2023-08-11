FROM golang:latest as build
WORKDIR /home/
COPY . .
RUN go build ./cmd/main.go

FROM ubuntu:22.04 as production
COPY --from=build /home/main .
CMD ["./main"]